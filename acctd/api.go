package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	mb "github.com/letterj/oohhc/proto/account"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

// Payload ...
type Payload struct {
	ID         string
	Name       string
	Token      string
	Status     string
	CreateDate int64
	DeleteDate int64
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
	var status string
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return &mb.CreateAcctResponse{Status: "Invalid Credintials"}, errors.New("Permission Denied")
	}
	// validate account string
	// validate new account does not exit
	//TODO:

	// create account information
	// Group:		/acct
	// Member:  "(uuid)"
	// Value:   { "id": "uuid", "name": "name", "apikey": "12345",
	//            "status": "active", "createdate": <timestamp>,
	//            "deletedate": <timestamp> }
	var p Payload
	group := "/acct"
	member := uuid.NewV4().String()
	// build payload
	p.ID = member
	p.Name = r.Acct
	p.Token = uuid.NewV4().String()
	p.Status = "active"
	p.CreateDate = time.Now().Unix()
	p.DeleteDate = 0
	detail, err := json.Marshal(p)
	if err != nil {
		status = fmt.Sprintf("account %s was not created", r.Acct)
		return &mb.CreateAcctResponse{Status: status}, err
	}
	// write information into the group store
	result, err := s.acctws.writeGStore(group, member, detail)
	if err != nil {
		status = fmt.Sprintf("account %s was not created", r.Acct)
		return &mb.CreateAcctResponse{Status: status}, err
	}
	log.Println(result)
	status = fmt.Sprintf("account %s was created with id %s", r.Acct, member)
	return &mb.CreateAcctResponse{Status: status}, nil
}

// ListAcct ...
func (s *AccountAPIServer) ListAcct(ctx context.Context, r *mb.ListAcctRequest) (*mb.ListAcctResponse, error) {
	var status string
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return &mb.ListAcctResponse{Account: nil, Status: "Invalid Credintials"}, errors.New("Permission Denied")
	}
	// validate account string
	// build the group store request
	group := "/acct"
	fmt.Println("group", group)

	// try and get account details form the group store
	result, err := s.acctws.lookupGStore(group)
	if err != nil {
		status = fmt.Sprintf("Problem looking up accounts %s", "/acct")
		return &mb.ListAcctResponse{Account: nil, Status: status}, err
	}
	// get information from the group store
	fmt.Println(result)

	status = "TODO:  A listing of all accounts"
	return &mb.ListAcctResponse{Account: nil, Status: status}, nil
}

// ShowAcct ...
func (s *AccountAPIServer) ShowAcct(ctx context.Context, r *mb.ShowAcctRequest) (*mb.ShowAcctResponse, error) {
	var status string
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return &mb.ShowAcctResponse{Account: nil, Status: "Invalid Credintials"}, errors.New("Permission Denied")
	}
	// validate account string

	// build the group store request
	group := "/acct"
	member := r.Acct
	fmt.Println("group", group)
	fmt.Println("member", member)

	// try and get account details form the group store
	result, err := s.acctws.getGStore(group, member)
	if err != nil {
		status = fmt.Sprintf("Problem looking up %s.", r.Acct)
		return &mb.ShowAcctResponse{Account: nil, Status: status}, err
	}
	if result == "" {
		status = fmt.Sprintf("Account %s not found", r.Acct)
		return &mb.ShowAcctResponse{Account: nil, Status: status}, errors.New("Not Found")
	}
	var p Payload
	err = json.Unmarshal([]byte(result), &p)
	if err != nil {
		status = fmt.Sprintf("Details parsing error for %s", r.Acct)
		return &mb.ShowAcctResponse{Account: nil, Status: status}, err
	}
	a := &mb.Account{
		Id:         p.ID,
		Name:       p.Name,
		Apikey:     p.Token,
		Status:     p.Status,
		CreateDate: p.CreateDate,
		DeleteDate: p.DeleteDate,
	}
	log.Println(result)
	status = fmt.Sprintf("account %s was found with id %s", r.Acct, member)
	return &mb.ShowAcctResponse{Account: a, Status: status}, nil
}

// DeleteAcct ...
func (s *AccountAPIServer) DeleteAcct(ctx context.Context, r *mb.DeleteAcctRequest) (*mb.DeleteAcctResponse, error) {
	var status string
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return &mb.DeleteAcctResponse{Status: "Invalid Credintials"}, errors.New("Permission Denied")
	}
	// validate account string
	// get information from the group store
	group := "/acct"
	member := r.Acct
	fmt.Println("group", group)
	fmt.Println("member", member)

	// try and get account details form the group store
	result, err := s.acctws.getGStore(group, member)
	if err != nil {
		status = fmt.Sprintf("Problem looking up %s.", r.Acct)
		return &mb.DeleteAcctResponse{Status: status}, err
	}
	if result == "" {
		status = fmt.Sprintf("Account %s not found", r.Acct)
		return &mb.DeleteAcctResponse{Status: status}, errors.New("Not Found")
	}
	var p Payload
	err = json.Unmarshal([]byte(result), &p)
	if err != nil {
		status = fmt.Sprintf("Details parsing error for %s", r.Acct)
		return &mb.DeleteAcctResponse{Status: status}, err
	}
	// only active accounts can be marked as deleted
	if p.Status != "active" || p.DeleteDate != 0 {
		status = fmt.Sprintf("Account %s not in an active state. State: %s, Deleted Date: %d", r.Acct, p.Status, p.DeleteDate)
		return &mb.DeleteAcctResponse{Status: status}, errors.New("Incorrect account status.")
	}
	// send delete to the group store
	p.Status = "deleted"
	p.DeleteDate = time.Now().Unix()
	detail, err := json.Marshal(p)
	if err != nil {
		status = fmt.Sprintf("account %s was not created", r.Acct)
		return &mb.DeleteAcctResponse{Status: status}, err
	}
	// write updated information into the group store
	result, err = s.acctws.writeGStore(group, member, detail)
	if err != nil {
		status = fmt.Sprintf("account %s was not deleted", r.Acct)
		return &mb.DeleteAcctResponse{Status: status}, err
	}
	log.Println(result)
	status = fmt.Sprintf("account %s was deleted", r.Acct)
	return &mb.DeleteAcctResponse{Status: status}, nil
}

// UpdateAcct ...
func (s *AccountAPIServer) UpdateAcct(ctx context.Context, r *mb.UpdateAcctRequest) (*mb.UpdateAcctResponse, error) {
	var status string
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return &mb.UpdateAcctResponse{Account: nil, Status: "Invalid Credintials"}, errors.New("Permission Denied")
	}
	// validate account string
	// pull the account information

	// write new information to the group store

	status = "TODO an update"
	return &mb.UpdateAcctResponse{Account: nil, Status: status}, nil
}
