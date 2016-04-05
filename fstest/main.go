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
	app.Version = "0.0.1"
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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "T, token",
					Value: "",
					Usage: "Token",
				},
			},
			Action: func(c *cli.Context) {
				if c.Args().Present() {
					fmt.Println("Invalid syntax for create.")
					os.Exit(1)
				}
				if c.String("token") == "" {
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
				result, err := ws.ShowFS(context.Background(), &mb.ShowFSRequest{Acctnum: acctNum, FSid: fsNum, Token: c.String("token")})
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
