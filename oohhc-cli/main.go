package main

import (
	"crypto/tls"
	"log"
	"os"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
	mb "github.com/letterj/oohhc/proto/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// main control function for oohhc-cli
func main() {

	// Set the location of the oohhc-acctd server
	var serverAddr string
	envServerAddr := os.Getenv("OOHHC_ACCT_SERVER_ADDR")
	if envServerAddr != "" {
		serverAddr = envServerAddr
	} else {
		serverAddr = "127.0.0.1:8449"
	}

	// Setup grpc
	var opts []grpc.DialOption
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	// Connect to oohhc-acctd
	ws := mb.NewAccountApiClient(conn)

	// Process command line arguments
	var accessKey string
	var acctNum string

	app := cli.NewApp()
	app.Name = "oohhc-cli"
	app.Usage = "Client used to manage accounts for FSAAAS"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "key, k",
			Value:       "",
			Usage:       "Access key for oohhc-acctd",
			EnvVar:      "OOHHC_ACCESS_KEY",
			Destination: &accessKey,
		},
		cli.StringFlag{
			Name:   "server, s",
			Value:  "127.0.0.1:8449",
			Usage:  "Address of the oohhc-acctd server",
			EnvVar: "OOHHC_SERVER_ADDRESS",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create a new account",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, N",
					Value: "",
					Usage: "Name for an account.",
				},
			},
			Action: func(c *cli.Context) {
				acctName := c.String("name")
				if !validAcctName(acctName) {
					log.Fatalf("Invalid Account String: %q", acctName)
					os.Exit(1)
				}
				result, err := ws.CreateAcct(context.Background(), &mb.CreateAcctRequest{Acctname: acctName, Superkey: accessKey})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					os.Exit(1)
				}
				log.Printf("CREATE Result: %s", result.Status)
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list all accounts",
			Action: func(c *cli.Context) {
				result, err := ws.ListAcct(context.Background(), &mb.ListAcctRequest{Superkey: accessKey})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					os.Exit(1)
				}
				log.Printf("Result: %s\n", result.Status)
				log.Printf("LIST Results: %s", result.Payload)
			},
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "details on a specific account",
			Action: func(c *cli.Context) {
				acctNum = c.Args().First()
				result, err := ws.ShowAcct(context.Background(), &mb.ShowAcctRequest{Acctnum: acctNum, Superkey: accessKey})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					os.Exit(1)
				}
				log.Printf("GET Result: %s\n", result.Status)
				log.Printf("Account: %s", result.Payload)
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "mark an account deleted",
			Action: func(c *cli.Context) {
				acctNum = c.Args().First()
				result, err := ws.DeleteAcct(context.Background(), &mb.DeleteAcctRequest{Acctnum: acctNum, Superkey: accessKey})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					os.Exit(1)
				}
				log.Printf("DELETE Result: %s\n", result.Status)
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update the information on an account",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, N",
					Value: "",
					Usage: "New name for an account.",
				},
				cli.StringFlag{
					Name:  "token, T",
					Value: "",
					Usage: "New token for the account",
				},
				cli.StringFlag{
					Name:  "status, S",
					Value: "",
					Usage: "New status for an account.",
				},
			},
			Action: func(c *cli.Context) {
				acctNum = c.Args().First()
				if c.String("name") != "" && !validAcctName(c.String("name")) {
					log.Fatalf("Invalid Account String: %q", c.String("name"))
					os.Exit(1)
				}
				modAcct := &mb.ModAccount{
					Name:   c.String("name"),
					Token:  c.String("token"),
					Status: c.String("status"),
				}
				result, err := ws.UpdateAcct(context.Background(), &mb.UpdateAcctRequest{Acctnum: acctNum, Superkey: accessKey, ModAcct: modAcct})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					os.Exit(1)
				}
				log.Printf("UPDATE Status Result: %s\n", result.Status)
				log.Printf("UPDATE Result: %s", result.Payload)
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
