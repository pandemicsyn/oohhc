package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/url"
	"os"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
	mb "github.com/letterj/oohhc/proto/filesystem"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// main control function for oohhc-cli
func main() {

	// Process command line arguments
	var token string
	var acctNum string
	var fsNum string
	var serverAddr string

	app := cli.NewApp()
	app.Name = "fstest"
	app.Usage = "Client used to test filesysd"
	app.Version = "0.0.2"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "token, T",
			Value:       "",
			Usage:       "Access token",
			EnvVar:      "OOHHC_TOKEN_KEY",
			Destination: &token,
		},
		cli.StringFlag{
			Name:        "server, s",
			Value:       "127.0.0.1:8448",
			Usage:       "Address of the oohhc-acctd server",
			EnvVar:      "OOHHC_FILESYS_SERVER_ADDR",
			Destination: &serverAddr,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "show",
			Usage: "Show a File Systems",
			Action: func(c *cli.Context) {
				if !c.Args().Present() {
					fmt.Println("Invalid syntax for show.")
					os.Exit(1)
				}
				if token == "" {
					fmt.Println("Token is required")
				}
				u, err := url.Parse(c.Args().Get(0))
				if err != nil {
					panic(err)
				}
				fmt.Println(u.Scheme)
				acctNum = u.Host
				fsNum = u.Path[1:]
				conn := setupWS(serverAddr)
				ws := mb.NewFileSystemAPIClient(conn)
				result, err := ws.ShowFS(context.Background(), &mb.ShowFSRequest{Acctnum: acctNum, FSid: fsNum, Token: token})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					conn.Close()
					os.Exit(1)
				}
				conn.Close()
				log.Printf("Result: %s\n", result.Status)
				log.Printf("SHOW Results: %s", result.Payload)
			},
		},
		{
			Name:  "create",
			Usage: "Create a File Systems",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, N",
					Value: "",
					Usage: "Name of the file system",
				},
			},
			Action: func(c *cli.Context) {
				if !c.Args().Present() {
					fmt.Println("Invalid syntax for show.")
					os.Exit(1)
				}
				if token == "" {
					fmt.Println("Token is required")
				}
				u, err := url.Parse(c.Args().Get(0))
				if err != nil {
					fmt.Printf("Url parse error: %v", err)
					os.Exit(1)
				}
				fmt.Println(u.Scheme)
				acctNum = u.Host
				if u.Path != "" {
					fmt.Println("Invalid url scheme")
					os.Exit(1)
				}
				if c.String("name") == "" {
					fmt.Println("File system name is a required field.")
					os.Exit(1)
				}
				conn := setupWS(serverAddr)
				ws := mb.NewFileSystemAPIClient(conn)
				result, err := ws.CreateFS(context.Background(), &mb.CreateFSRequest{Acctnum: acctNum, FSName: c.String("name"), Token: token})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					conn.Close()
					os.Exit(1)
				}
				conn.Close()
				log.Printf("Result: %s\n", result.Status)
				log.Printf("Create Results: %s", result.Payload)
			},
		},
		{
			Name:  "list",
			Usage: "List File Systems for an account",
			Action: func(c *cli.Context) {
				if !c.Args().Present() {
					fmt.Println("Invalid syntax for list.")
					os.Exit(1)
				}
				if token == "" {
					fmt.Println("Token is required")
				}
				u, err := url.Parse(c.Args().Get(0))
				if err != nil {
					fmt.Println("Invalid url scheme")
					os.Exit(1)
				}
				fmt.Println(u.Scheme)
				acctNum = u.Host
				if u.Path != "" {
					fmt.Println("Invaid url")
					os.Exit(1)
				}
				conn := setupWS(serverAddr)
				ws := mb.NewFileSystemAPIClient(conn)
				result, err := ws.ListFS(context.Background(), &mb.ListFSRequest{Acctnum: acctNum, Token: token})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					conn.Close()
					os.Exit(1)
				}
				conn.Close()
				log.Printf("Result: %s\n", result.Status)
				log.Printf("LIST Results: %s", result.Payload)
			},
		},
		{
			Name:  "delete",
			Usage: "Delete a File Systems",
			Action: func(c *cli.Context) {
				if !c.Args().Present() {
					fmt.Println("Invalid syntax for delete.")
					os.Exit(1)
				}
				if token == "" {
					fmt.Println("Token is required")
				}
				u, err := url.Parse(c.Args().Get(0))
				if err != nil {
					fmt.Println("Invalid url scheme")
					os.Exit(1)
				}
				fmt.Println(u.Scheme)
				acctNum = u.Host
				fsNum = u.Path[1:]
				conn := setupWS(serverAddr)
				ws := mb.NewFileSystemAPIClient(conn)
				result, err := ws.DeleteFS(context.Background(), &mb.DeleteFSRequest{Acctnum: acctNum, FSid: fsNum, Token: token})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					conn.Close()
					os.Exit(1)
				}
				conn.Close()
				log.Printf("Result: %s\n", result.Status)
				log.Printf("Delete Results: %s", result.Payload)
			},
		},
		{
			Name:  "update",
			Usage: "Update a File Systems",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, N",
					Value: "",
					Usage: "Name of the file system",
				},
				cli.StringFlag{
					Name:  "S, status",
					Value: "",
					Usage: "Status of the file system",
				},
			},
			Action: func(c *cli.Context) {
				if !c.Args().Present() {
					fmt.Println("Invalid syntax for delete.")
					os.Exit(1)
				}
				if token == "" {
					fmt.Println("Token is required")
				}
				u, err := url.Parse(c.Args().Get(0))
				if err != nil {
					fmt.Printf("Url Parse error: %v", err)
					os.Exit(1)
				}
				fmt.Println(u.Scheme)
				acctNum = u.Host
				fsNum = u.Path[1:]
				if c.String("name") != "" && !validAcctName(c.String("name")) {
					fmt.Printf("Invalid File System String: %q\n", c.String("name"))
					os.Exit(1)
				}
				fsMod := &mb.ModFS{
					Name:   c.String("name"),
					Status: c.String("status"),
				}
				conn := setupWS(serverAddr)
				ws := mb.NewFileSystemAPIClient(conn)
				result, err := ws.UpdateFS(context.Background(), &mb.UpdateFSRequest{Acctnum: acctNum, FSid: fsNum, Token: token, Filesys: fsMod})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					conn.Close()
					os.Exit(1)
				}
				conn.Close()
				log.Printf("Result: %s\n", result.Status)
				log.Printf("Update Results: %s", result.Payload)
			},
		},
		{
			Name:  "grant",
			Usage: "Grant an Addr access to a File Systems",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "addr",
					Value: "",
					Usage: "Address to Grant",
				},
			},
			Action: func(c *cli.Context) {
				if !c.Args().Present() {
					fmt.Println("Invalid syntax for delete.")
					os.Exit(1)
				}
				if token == "" {
					fmt.Println("Token is required")
					os.Exit(1)
				}
				if c.String("addr") == "" {
					fmt.Println("addr is required")
					os.Exit(1)
				}
				u, err := url.Parse(c.Args().Get(0))
				if err != nil {
					fmt.Println("Invalid url scheme")
					os.Exit(1)
				}
				fmt.Println(u.Scheme)
				acctNum = u.Host
				fsNum = u.Path[1:]
				conn := setupWS(serverAddr)
				ws := mb.NewFileSystemAPIClient(conn)
				result, err := ws.GrantAddrFS(context.Background(), &mb.GrantAddrFSRequest{Acctnum: acctNum, FSid: fsNum, Token: token, Addr: c.String("addr")})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					conn.Close()
					os.Exit(1)
				}
				conn.Close()
				log.Printf("Result: %s\n", result.Status)
			},
		},
		{
			Name:  "revoke",
			Usage: "Revoke an Addr's access to a File Systems",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "addr",
					Value: "",
					Usage: "Address to Revoke",
				},
			},
			Action: func(c *cli.Context) {
				if !c.Args().Present() {
					fmt.Println("Invalid syntax for revoke.")
					os.Exit(1)
				}
				if token == "" {
					fmt.Println("Token is required")
					os.Exit(1)
				}
				if c.String("addr") == "" {
					fmt.Println("addr is required")
					os.Exit(1)
				}
				u, err := url.Parse(c.Args().Get(0))
				if err != nil {
					fmt.Println("Invalid url scheme")
					os.Exit(1)
				}
				fmt.Println(u.Scheme)
				acctNum = u.Host
				fsNum = u.Path[1:]
				conn := setupWS(serverAddr)
				ws := mb.NewFileSystemAPIClient(conn)
				result, err := ws.RevokeAddrFS(context.Background(), &mb.RevokeAddrFSRequest{Acctnum: acctNum, FSid: fsNum, Token: token, Addr: c.String("addr")})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					conn.Close()
					os.Exit(1)
				}
				conn.Close()
				log.Printf("Result: %s\n", result.Status)
			},
		},
		{
			Name:  "verify",
			Usage: "Verify an Addr has access to a file system",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "addr",
					Value: "",
					Usage: "Address to check",
				},
			},
			Action: func(c *cli.Context) {
				if !c.Args().Present() {
					fmt.Println("Invalid syntax for revoke.")
					os.Exit(1)
				}
				if c.String("addr") == "" {
					fmt.Println("addr is required")
					os.Exit(1)
				}
				u, err := url.Parse(c.Args().Get(0))
				if err != nil {
					fmt.Println("Invalid url scheme")
					os.Exit(1)
				}
				fmt.Println(u.Scheme)
				fsNum = u.Host
				conn := setupWS(serverAddr)
				ws := mb.NewFileSystemAPIClient(conn)
				result, err := ws.LookupAddrFS(context.Background(), &mb.LookupAddrFSRequest{FSid: fsNum, Addr: c.String("addr")})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					conn.Close()
					os.Exit(1)
				}
				conn.Close()
				log.Printf("Result: %s\n", result.Status)
			},
		},
	}
	app.Run(os.Args)
}

// Validate the account string passed in from the command line
func validAcctName(a string) bool {
	//TODO: Determine what needs to be done to validate
	return true
}

func setupWS(svr string) *grpc.ClientConn {
	var opts []grpc.DialOption
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	opts = append(opts, grpc.WithTransportCredentials(creds))
	c, err := grpc.Dial(svr, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	return c

}
