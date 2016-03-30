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
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	fb "github.com/letterj/oohhc/proto/filesystem"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
)

var errf = grpc.Errorf

// AcctPayLoad ... Account PayLoad
type AcctPayLoad struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Token      string `json:"token"`
	Status     string `json:"status"`
	CreateDate int64  `json:"createdate"`
	DeleteDate int64  `json:"deletedate"`
}

// FileSysPayLoad ... File System PayLoad
type FileSysPayLoad struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	SizeInBytes int64  `json:"sizeinbytes"`
	Status      string `json:"status"`
	CreateDate  int64  `json:"createdate"`
	DeleteDate  int64  `json:"deletedate"`
}

// AddrPayLoad ... IP Address Address PayLoad
type AddrPayLoad struct {
	ID         string `json:"id"`
	Addr       string `json:"addr"`
	Status     string `json:"status"`
	CreateDate int64  `json:"createdate"`
	DeleteDate int64  `json:"deletedate"`
}

// FileSystemAPIServer is used to implement oohhc
type FileSystemAPIServer struct {
	fsws *FileSystemWS
}

// NewFileSystemAPIServer ...
func NewFileSystemAPIServer(filesysWS *FileSystemWS) *FileSystemAPIServer {
	s := new(FileSystemAPIServer)
	s.fsws = filesysWS
	return s
}

// CreateFS ...
func (s *FileSystemAPIServer) CreateFS(ctx context.Context, r *fb.CreateFSRequest) (*fb.CreateFSResponse, error) {
	var status string
	var result string
	var acctData AcctPayLoad
	var err error
	// Get incomming ip
	pr, ok := peer.FromContext(ctx)
	if ok {
		fmt.Println(pr.Addr)
	}
	// getAcct data
	acctData, err = s.getAcct("/acct", r.Acctnum)
	if err != nil {
		log.Printf("Error %v on lookup for account %s", err, r.Acctnum)
		return nil, err
	}
	// validate token
	if acctData.Token != r.Token {
		return nil, errf(codes.PermissionDenied, "%s", "Invalid Token")
	}
	// Check for to see if file system name exists
	fs := fmt.Sprintf("/acct/%s/fs", acctData.ID)
	err = s.dupNameCheck(fs, r.FSName)
	if err != nil {
		log.Printf("Precondition Failed: %v\n...", err)
		return nil, errf(codes.FailedPrecondition, "%v", err)
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
	var acctData AcctPayLoad
	var err error
	// Get incomming ip
	pr, ok := peer.FromContext(ctx)
	if ok {
		fmt.Println(pr.Addr)
	}
	// getAcct data
	acctData, err = s.getAcct("/acct", r.Acctnum)
	if err != nil {
		log.Printf("Error %v on lookup for account %s", err, r.Acctnum)
		return nil, err
	}
	// validate token
	if acctData.Token != r.Token {
		return nil, errf(codes.PermissionDenied, "%s", "Invalid Token")
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
	var acctData AcctPayLoad
	var err error
	// Get incomming ip
	pr, ok := peer.FromContext(ctx)
	if ok {
		fmt.Println(pr.Addr)
	}
	// getAcct data
	acctData, err = s.getAcct("/acct", r.Acctnum)
	if err != nil {
		log.Printf("Error %v on lookup for account %s", err, r.Acctnum)
		return nil, err
	}
	// validate token
	if acctData.Token != r.Token {
		return nil, errf(codes.PermissionDenied, "%s", "Invalid Token")
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
	var acctData AcctPayLoad
	var err error
	// Get incomming ip
	pr, ok := peer.FromContext(ctx)
	if ok {
		fmt.Println(pr.Addr)
	}
	// getAcct data
	acctData, err = s.getAcct("/acct", r.Acctnum)
	if err != nil {
		log.Printf("Error %v on lookup for account %s", err, r.Acctnum)
		return nil, err
	}
	// validate token
	if acctData.Token != r.Token {
		return nil, errf(codes.PermissionDenied, "%s", "Invalid Token")
	}

	// DO STUFF

	// Prep things to return
	status = fmt.Sprintf("filesystem %s in account %s was deleted", r.FSName, r.Acctnum)
	result = "STUFF"
	return &fb.DeleteFSResponse{Payload: result, Status: status}, nil
}

// UpdateFS ...
func (s *FileSystemAPIServer) UpdateFS(ctx context.Context, r *fb.UpdateFSRequest) (*fb.UpdateFSResponse, error) {
	var status string
	var result string
	var acctData AcctPayLoad
	var err error
	// Get incomming ip
	pr, ok := peer.FromContext(ctx)
	if ok {
		fmt.Println(pr.Addr)
	}
	// getAcct data
	acctData, err = s.getAcct("/acct", r.Acctnum)
	if err != nil {
		log.Printf("Error %v on lookup for account %s", err, r.Acctnum)
		return nil, err
	}
	// validate token
	if acctData.Token != r.Token {
		return nil, errf(codes.PermissionDenied, "%s", "Invalid Token")
	}

	// DO stuff
	status = fmt.Sprintf("filesystem %s with id %s was updated", r.FSName, r.Acctnum)
	result = "STUFF"
	return &fb.UpdateFSResponse{Payload: result, Status: status}, nil
}

// GrantAddrFS ...
func (s *FileSystemAPIServer) GrantAddrFS(ctx context.Context, r *fb.GrantAddrFSRequest) (*fb.GrantAddrFSResponse, error) {
	var status string
	var err error
	var acctData AcctPayLoad
	var fsData FileSysPayLoad
	// Get incomming ip
	pr, ok := peer.FromContext(ctx)
	if ok {
		fmt.Println(pr.Addr)
	}
	// getAcct data
	acctData, err = s.getAcct("/acct", r.Acctnum)
	if err != nil {
		log.Printf("Error %v on lookup for account %s", err, r.Acctnum)
		return nil, err
	}
	// validate token
	if acctData.Token != r.Token {
		return nil, errf(codes.PermissionDenied, "%s", "Invalid Token")
	}
	// getFS data
	fs := fmt.Sprintf("/acct/%s/fs/%s", r.Acctnum, r.FSName)
	fsData, err = s.getFS(fs, r.Acctnum)
	if err != nil {
		log.Printf("Error %v on lookup for account %s", err, r.Acctnum)
		return nil, err
	}

	// Check for duplicate addresses
	fsaddr := fmt.Sprintf("/acct/%s/fs/%s/addr", acctData.ID, r.FSName)
	err = s.dupNameCheck(fsaddr, r.Addr)
	if err != nil {
		log.Printf("Precondition Failed: %v\n...", err)
		return nil, errf(codes.FailedPrecondition, "%v", err)
	}

	// DO stuff
	status = fmt.Sprintf("addr %s for filesystem %s with account id %s was granted", r.Addr, r.FSName, r.Acctnum)
	return &fb.GrantAddrFSResponse{Status: status}, nil
}

// RevokeAddrFS ...
func (s *FileSystemAPIServer) RevokeAddrFS(ctx context.Context, r *fb.RevokeAddrFSRequest) (*fb.RevokeAddrFSResponse, error) {
	var status string
	var err error
	var acctData AcctPayLoad
	// Get incomming ip
	pr, ok := peer.FromContext(ctx)
	if ok {
		fmt.Println(pr.Addr)
	}
	// getAcct data
	acctData, err = s.getAcct("/acct", r.Acctnum)
	if err != nil {
		log.Printf("Error %v on lookup for account %s", err, r.Acctnum)
		return nil, err
	}
	// validate token
	if acctData.Token != r.Token {
		return nil, errf(codes.PermissionDenied, "%s", "Invalid Token")
	}

	// DO stuff
	status = fmt.Sprintf("addr %s for filesystem %s with account id %s was revoked", r.Addr, r.FSName, r.Acctnum)
	return &fb.RevokeAddrFSResponse{Status: status}, nil
}

// getAcct ...
func (s *FileSystemAPIServer) getAcct(k string, ck string) (AcctPayLoad, error) {
	var a AcctPayLoad
	var err error

	data, err := s.fsws.getGStore(k, ck)
	if err != nil {
		return a, errf(codes.Internal, "%v", err)
	}
	if data == "" {
		return a, errf(codes.NotFound, "%s", "Account Not Found")
	}
	log.Printf("Got back %s", data)
	err = json.Unmarshal([]byte(data), &a)
	if err != nil {
		return a, errf(codes.Internal, "%v", err)
	}
	if a.Status != "active" {
		return a, errf(codes.NotFound, "%s", "Account Not Active")
	}
	return a, nil
}

// getAcct ...
func (s *FileSystemAPIServer) getFS(k string, ck string) (FileSysPayLoad, error) {
	var fs FileSysPayLoad
	var err error

	data, err := s.fsws.getGStore(k, ck)
	if err != nil {
		return fs, errf(codes.Internal, "%v", err)
	}
	if data == "" {
		return fs, errf(codes.NotFound, "%s", "FileSystem Not Found")
	}
	log.Printf("Got back %s", data)
	err = json.Unmarshal([]byte(data), &fs)
	if err != nil {
		return fs, errf(codes.Internal, "%v", err)
	}
	if fs.Status != "active" {
		return fs, errf(codes.NotFound, "%s", "Account Not Active")
	}
	return fs, nil
}

// dupNameCheck ...
func (s *FileSystemAPIServer) dupNameCheck(fs string, n string) error {
	var fsdata FileSysPayLoad
	// try and get file system details form the group store
	data, err := s.fsws.lookupGStore(fs)
	log.Printf("Data from the list: %v", data)
	if err != nil {
		log.Printf("Problem talking to Group Store: %v", err)
		return err
	}
	if data == "" {
		return nil
	}
	fsList := strings.Split(data, "|")
	log.Printf("Number of file systems in the list: %v", len(fsList))
	log.Printf("File System: %v", fsList)
	for i := 0; i < len(fsList); i++ {
		if fsList[i] != "" {
			err = json.Unmarshal([]byte(fsList[i]), &fsdata)
			if err != nil {
				log.Printf("Unmarshal Error: %v", err)
				return err
			}
			if fsdata.Status == "active" {
				if strings.ToLower(fsdata.Name) == strings.ToLower(n) {
					log.Printf("FileSystem Name already exists: %s", n)
					return errors.New("Name already exists")
				}
			}
		}
	}
	return nil
}

// dupAddrCheck ...
func (s *FileSystemAPIServer) dupAddrCheck(fsaddr string, addr string) error {
	var addrdata AddrPayLoad
	// try and get addr details form the group store
	data, err := s.fsws.lookupGStore(fsaddr)
	log.Printf("Data from the list: %v", data)
	if err != nil {
		log.Printf("Problem talking to Group Store: %v", err)
		return err
	}
	if data == "" {
		return nil
	}
	fsAddrList := strings.Split(data, "|")
	log.Printf("Number of addresses in the list: %v", len(fsAddrList))
	log.Printf("Address: %v", fsAddrList)
	for i := 0; i < len(fsAddrList); i++ {
		if fsAddrList[i] != "" {
			err = json.Unmarshal([]byte(fsAddrList[i]), &addrdata)
			if err != nil {
				log.Printf("Unmarshal Error: %v", err)
				return err
			}
			if addrdata.Status == "active" {
				if addrdata.Addr == addr {
					log.Printf("Addr already exists: %s", addr)
					return errors.New("Addr already exists")
				}
			}
		}
	}
	return nil
}
