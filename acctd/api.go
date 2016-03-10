package main

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	mb "github.com/letterj/oohhc/proto/account"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
)

var errf = grpc.Errorf

// PayLoad ...
type PayLoad struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Token      string `json:"token"`
	Status     string `json:"status"`
	CreateDate int64  `json:"createdate"`
	DeleteDate int64  `json:"deletedate"`
}

// AccountAPIServer is used to implement grpchello.CfsAdminApiServer
type AccountAPIServer struct {
	acctws *AccountWS
}

// NewAccountAPIServer ...
func NewAccountAPIServer(acctws *AccountWS) *AccountAPIServer {
	s := new(AccountAPIServer)
	s.acctws = acctws
	return s
}

// CreateAcct ...
func (s *AccountAPIServer) CreateAcct(ctx context.Context, r *mb.CreateAcctRequest) (*mb.CreateAcctResponse, error) {
	// Verify client is from 127.0.0.1
	pr, _ := peer.FromContext(ctx)
	if pr.Addr.String() != "127.0.0.1" {
		log.Printf("Invalid Access attempt from %s", pr.Addr.String())
		return nil, errf(codes.Canceled, "%s", "oohhc-acctd can only be accessed locally")
	}
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return nil, errf(codes.PermissionDenied, "%s", "Invalid Key")
	}

	// validate new account does not exit
	err := s.duplicateName(r.Acctname)
	if err != nil {
		return nil, errf(codes.FailedPrecondition, "%v", err)
	}

	// create account information
	// Group:		/acct
	// Member:  "(uuid)"
	// Value:   { "id": "uuid", "name": "name", "apikey": "12345",
	//            "status": "active", "createdate": <timestamp>,
	//            "deletedate": <timestamp> }
	var p PayLoad
	g := "/acct"
	m := uuid.NewV4().String()
	// build payload
	p.ID = m
	p.Name = r.Acctname
	p.Token = uuid.NewV4().String()
	p.Status = "active"
	p.CreateDate = time.Now().Unix()
	p.DeleteDate = 0
	d, err := json.Marshal(p)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	// write information into the group store
	_, err = s.acctws.writeGStore(g, m, d)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	return &mb.CreateAcctResponse{Status: m}, nil
}

// ListAcct ...
func (s *AccountAPIServer) ListAcct(ctx context.Context, r *mb.ListAcctRequest) (*mb.ListAcctResponse, error) {
	// Verify client is from 127.0.0.1
	pr, _ := peer.FromContext(ctx)
	if pr.Addr.String() != "127.0.0.1" {
		log.Printf("Invalid Access attempt from %s", pr.Addr.String())
		return nil, errf(codes.Canceled, "%s", "oohhc-acctd can only be accessed locally")
	}
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return nil, errf(codes.PermissionDenied, "%s", "Invalid Key")
	}
	// build the group store request
	g := "/acct"

	// try and get account details form the group store
	data, err := s.acctws.lookupGStore(g)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	return &mb.ListAcctResponse{Payload: data, Status: "OK"}, nil
}

// ShowAcct ...
func (s *AccountAPIServer) ShowAcct(ctx context.Context, r *mb.ShowAcctRequest) (*mb.ShowAcctResponse, error) {
	// Verify client is from 127.0.0.1
	pr, _ := peer.FromContext(ctx)
	if pr.Addr.String() != "127.0.0.1" {
		log.Printf("Invalid Access attempt from %s", pr.Addr.String())
		return nil, errf(codes.Canceled, "%s", "oohhc-acctd can only be accessed locally")
	}
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return nil, errf(codes.PermissionDenied, "%s", "Invalid Key")
	}

	// build the group store request
	g := "/acct"
	m := r.Acctnum

	// try and get account details form the group store
	data, err := s.acctws.getGStore(g, m)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	if data == "" {
		return nil, errf(codes.NotFound, "%s", "Account Not Found")
	}
	var p PayLoad
	err = json.Unmarshal([]byte(data), &p)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	return &mb.ShowAcctResponse{Payload: data, Status: "OK"}, nil
}

// DeleteAcct ...
func (s *AccountAPIServer) DeleteAcct(ctx context.Context, r *mb.DeleteAcctRequest) (*mb.DeleteAcctResponse, error) {
	// Verify client is from 127.0.0.1
	pr, _ := peer.FromContext(ctx)
	if pr.Addr.String() != "127.0.0.1" {
		log.Printf("Invalid Access attempt from %s", pr.Addr.String())
		return nil, errf(codes.Canceled, "%s", "oohhc-acctd can only be accessed locally")
	}
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return nil, errf(codes.PermissionDenied, "%s", "Invalid Key")
	}
	// get information from the group store
	g := "/acct"
	m := r.Acctnum

	// try and get account details form the group store
	result, err := s.acctws.getGStore(g, m)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	if result == "" {
		return nil, errf(codes.NotFound, "%s", "Account Not Found")
	}
	var p PayLoad
	err = json.Unmarshal([]byte(result), &p)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	// only active accounts can be marked as deleted
	if p.Status != "active" || p.DeleteDate != 0 {
		return nil, errf(codes.InvalidArgument, "%s", "Passing Account Status")
	}
	// send delete to the group store
	p.Status = "deleted"
	p.DeleteDate = time.Now().Unix()
	d, err := json.Marshal(p)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	// write updated information into the group store
	result, err = s.acctws.writeGStore(g, m, d)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	return &mb.DeleteAcctResponse{Status: m}, nil
}

// UpdateAcct ...
func (s *AccountAPIServer) UpdateAcct(ctx context.Context, r *mb.UpdateAcctRequest) (*mb.UpdateAcctResponse, error) {
	// Verify client is from 127.0.0.1
	pr, _ := peer.FromContext(ctx)
	if pr.Addr.String() != "127.0.0.1" {
		log.Printf("Invalid Access attempt from %s", pr.Addr.String())
		return nil, errf(codes.Canceled, "%s", "oohhc-acctd can only be accessed locally")
	}
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return nil, errf(codes.PermissionDenied, "%s", "Invalid Key")
	}

	g := "/acct"
	m := r.Acctnum

	// try and get account details form the group store
	result, err := s.acctws.getGStore(g, m)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	if result == "" {
		return nil, errf(codes.NotFound, "%s", "Account Not Found")
	}
	var p PayLoad
	err = json.Unmarshal([]byte(result), &p)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	// update account information
	if r.ModAcct.Name != "" {
		// Check for duplicate Name
		err = s.duplicateName(r.ModAcct.Name)
		if err != nil {
			return nil, errf(codes.FailedPrecondition, "%v", err)
		}
		p.Name = r.ModAcct.Name
	}
	if r.ModAcct.Status != "" {
		if p.Status == "deleted" && r.ModAcct.Status != "deleted" {
			p.DeleteDate = 0
		}
		p.Status = r.ModAcct.Status
	}
	if r.ModAcct.Token != "" {
		p.Token = r.ModAcct.Token
	}
	// write new information to the group store
	d, err := json.Marshal(p)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	// write information into the group store
	_, err = s.acctws.writeGStore(g, m, d)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	// Pull updated data
	uresult, err := s.acctws.getGStore(g, m)
	if err != nil {
		return nil, errf(codes.Internal, "%v", err)
	}
	if uresult == "" {
		return nil, errf(codes.NotFound, "%s", "Update Not Found")
	}
	// Good request return
	return &mb.UpdateAcctResponse{Payload: uresult, Status: "OK"}, nil
}

// duplicateName will check to see if an account name already exists
func (s *AccountAPIServer) duplicateName(acctName string) error {
	var p PayLoad
	g := "/acct"
	// try and get account details form the group store
	data, err := s.acctws.lookupGStore(g)
	if err != nil {
		// figure out something to do if
		log.Printf("Problem talking to Group Store: %v", err)
		return err
	}
	aList := strings.Split(data, ",")
	for i := 0; i < len(aList); i++ {
		err = json.Unmarshal([]byte(aList[i]), &p)
		if err != nil {
			log.Printf("Unmarshal Error: %v", err)
			return err
		}
		if strings.ToLower(p.Name) == strings.ToLower(acctName) {
			log.Printf("Account Name already exists: %s", acctName)
			return errors.New("Account Name Exists")
		}
	}
	return nil
}
