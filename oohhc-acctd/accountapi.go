package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	mb "github.com/letterj/oohhc/proto/account"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

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

	// create account information
	// Group:		/acct
	// Member:  "(uuid)"
	// Value:   { "id": "uuid", "name": "name", "apikey": "12345",
	//            "status": "active", "createdate": <timestamp>,
	//            "deletedate": <timestamp> }
	group := "/acct"
	member := uuid.NewV4().String()
	token := uuid.NewV4().String()
	createDate := time.Now().Unix()
	deleteDate := 0
	details := fmt.Sprintf(`{"id": %s, "name": %s, "token": %s, "status": %s, "createdate": %d, "deletedate": %d}`, member, r.Acct, token, "active", createDate, deleteDate)

	// write information into the group store
	result, err := s.acctws.writeGStore(group, member, details)
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

	// get information from the group store

	status = "TODO"
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
		status = fmt.Sprintf("account %s was not found", r.Acct)
		return &mb.ShowAcctResponse{Account: nil, Status: status}, err
	}
	log.Println(result)
	status = fmt.Sprintf("account %s was found with id %s and values of %s", r.Acct, member, result)
	return &mb.ShowAcctResponse{Account: nil, Status: status}, nil
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

	// send delete to the group store

	status = "TODO a delete"
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
