// Structures used in Group Store
//  File System
//  /acct/(uuid)/fs  "(uuid)"    { "id": "uuid", "name": "name", "status": "active",
//                                "createdate": <timestamp>, "deletedate": <timestamp>
//                               }
//
// IP Address
// /acct/(uuid)/fs/(uuid)/addr "(uuid)"   { "id": uuid, "addr": "111.111.111.111", "status": "active",
//                                         "createdate": <timestamp>, "deletedate": <timestamp>
//                                       }

package main

import (
	"errors"
	"fmt"

	fb "github.com/letterj/oohhc/proto/filesystem"

	"golang.org/x/net/context"
)

// PayLoad ...
type PayLoad struct {
	ID         string
	Name       string
	Token      string
	Status     string
	CreateDate int64
	DeleteDate int64
}

// FileSystemAPIServer is used to implement oohhc
type FileSystemAPIServer struct {
	filesysWS *FileSystemWS
}

// NewFileSystemAPIServer ...
func NewFileSystemAPIServer(filesysWS *FileSystemWS) *FileSystemAPIServer {
	s := new(FileSystemAPIServer)
	s.filesysWS = filesysWS
	return s
}

// CreateFS ...
func (s *FileSystemAPIServer) CreateFS(ctx context.Context, r *fb.CreateFSRequest) (*fb.CreateFSResponse, error) {
	var status string
	var result string
	// validate token
	if !validToken(r.Acctnum, r.Token) {
		return &fb.CreateFSResponse{Status: "Invalid Credintials"}, errors.New("Permission Denied")
	}
	// DO STUFF

	// Prep things to return
	status = fmt.Sprintf("CREATE: %s  OK", r.Acctnum)
	result = fmt.Sprint("STUFF")
	return &fb.CreateFSResponse{Payload: result, Status: status}, nil
}

// ListFS ...
func (s *FileSystemAPIServer) ListFS(ctx context.Context, r *fb.ListFSRequest) (*fb.ListFSResponse, error) {
	var status string
	var result string
	// validate token
	if !validToken(r.Acctnum, r.Token) {
		return &fb.ListFSResponse{Status: "Invalid Credintials"}, errors.New("Permission Denied")
	}

	// DO STUFF

	// Prep things to return
	status = fmt.Sprintf("Complete file system listing for account %s", r.Acctnum)
	result = fmt.Sprintf("STUFF")
	return &fb.ListFSResponse{Payload: result, Status: status}, nil
}

// ShowFS ...
func (s *FileSystemAPIServer) ShowFS(ctx context.Context, r *fb.ShowFSRequest) (*fb.ShowFSResponse, error) {
	var status string
	var result string
	// validate token
	if !validToken(r.Acctnum, r.Token) {
		return &fb.ShowFSResponse{Payload: "", Status: "Invalid Credintials"}, errors.New("Permission Denied")
	}

	// DO STUFF

	// Prep things to return
	status = fmt.Sprintf("Complete file system listing for account %s", r.Acctnum)
	result = fmt.Sprintf("STUFF")
	return &fb.ShowFSResponse{Payload: result, Status: status}, nil
}

// DeleteFS ...
func (s *FileSystemAPIServer) DeleteFS(ctx context.Context, r *fb.DeleteFSRequest) (*fb.DeleteFSResponse, error) {
	var status string
	var result string
	// validate Token
	if !validToken(r.Acctnum, r.Token) {
		return &fb.DeleteFSResponse{Payload: "", Status: "Invalid Credintials"}, errors.New("Permission Denied")
	}

	// DO STUFF

	// Prep things to return
	status = fmt.Sprintf("filesystem %s in account %s was deleted", r.Fsname, r.Acctnum)
	result = "STUFF"
	return &fb.DeleteFSResponse{Payload: result, Status: status}, nil
}

// UpdateFS ...
func (s *FileSystemAPIServer) UpdateFS(ctx context.Context, r *fb.UpdateFSRequest) (*fb.UpdateFSResponse, error) {
	var status string
	var result string
	// validate token
	if !validToken(r.Acctnum, r.Token) {
		return &fb.UpdateFSResponse{Payload: "", Status: "Invalid Credintials"}, errors.New("Permission Denied")
	}

	// DO stuff
	status = fmt.Sprintf("filesystem %s with id %s was updated", r.Fsname, r.Acctnum)
	result = "STUFF"
	return &fb.UpdateFSResponse{Payload: result, Status: status}, nil
}

// GrantAddrFS ...
func (s *FileSystemAPIServer) GrantAddrFS(ctx context.Context, r *fb.GrantAddrFSRequest) (*fb.GrantAddrFSResponse, error) {
	var status string
	// validate token

	// DO stuff
	status = fmt.Sprintf("addr %s for filesystem %s with account id %s was granted", r.Addr, r.Fsname, r.Acctnum)
	return &fb.GrantAddrFSResponse{Status: status}, nil
}

// RevokeAddrFS ...
func (s *FileSystemAPIServer) RevokeAddrFS(ctx context.Context, r *fb.RevokeAddrFSRequest) (*fb.RevokeAddrFSResponse, error) {
	var status string
	// validate token

	// DO stuff
	status = fmt.Sprintf("addr %s for filesystem %s with account id %s was revoked", r.Addr, r.Fsname, r.Acctnum)
	return &fb.RevokeAddrFSResponse{Status: status}, nil
}

// validToken
func validToken(a string, t string) bool {
	// call the group store to get the info
	return true
}
