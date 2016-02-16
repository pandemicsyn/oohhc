package main

import (
	"crypto/tls"
	"fmt"
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
	var acctStr string

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
			Action: func(c *cli.Context) {
				acctStr = c.Args().First()
				if !validAcctStr(acctStr) {
					log.Fatalf("Invalid Account String: %q", acctStr)
					os.Exit(1)
				}
				result, err := ws.CreateAcct(context.Background(), &mb.CreateAcctRequest{Acct: acctStr, Superkey: accessKey})
				if err != nil {
					fmt.Println("key", accessKey)
					log.Fatalf("Bad Request: %v", err)
					os.Exit(1)
				}
				log.Printf("Result: %s", result.Status)
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list all accounts",
			Action: func(c *cli.Context) {
				result, err := ws.ListAcct(context.Background(), &mb.ListAcctRequest{Superkey: accessKey})
				if err != nil {
					fmt.Println("key", accessKey)
					log.Fatalf("Bad Request: %v", err)
					os.Exit(1)
				}
				log.Printf("Result: %s\n", result.Status)
				log.Printf("Result: %s", result.Account)
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "mark an account deleted",
			Action: func(c *cli.Context) {
				acctStr = c.Args().First()
				if !validAcctStr(acctStr) {
					log.Fatalf("Invalid Account String: %q", acctStr)
					os.Exit(1)
				}
				result, err := ws.DeleteAcct(context.Background(), &mb.DeleteAcctRequest{Acct: acctStr, Superkey: accessKey})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					os.Exit(1)
				}
				log.Printf("Result: %s\n", result.Status)
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
					Name:  "apikey, K",
					Value: "",
					Usage: "New apikey for the account",
				},
				cli.StringFlag{
					Name:  "status, S",
					Value: "",
					Usage: "New status for an account.",
				},
				cli.BoolFlag{
					Name:  "undelete, D",
					Usage: "Undelete an account",
				},
			},
			Action: func(c *cli.Context) {
				acctStr = c.Args().First()
				if !validAcctStr(acctStr) {
					log.Fatalf("Invalid Account String: %q", acctStr)
					os.Exit(1)
				}
				var newDate int64
				if c.Bool("undelete") {
					newDate = 0
				}
				modAcct := &mb.ModAccount{
					Name:       c.String("name"),
					Apikey:     c.String("apikey"),
					Status:     c.String("status"),
					Deletedate: newDate,
				}
				result, err := ws.UpdateAcct(context.Background(), &mb.UpdateAcctRequest{Acct: acctStr, Superkey: accessKey, ModAcct: modAcct})
				if err != nil {
					log.Fatalf("Bad Request: %v", err)
					os.Exit(1)
				}
				log.Printf("Result: %s\n", result.Status)
			},
		},
	}
	app.Run(os.Args)
}

// Validate the account string passed in from the command line
func validAcctStr(a string) bool {
	//TODO: Determine what needs to be done to validate
	return true
}
