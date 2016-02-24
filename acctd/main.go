package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	mb "github.com/letterj/oohhc/proto/account"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"net"
)

var (
	usetls             = flag.Bool("tls", true, "Connection uses TLS if true, else plain TCP")
	certFile           = flag.String("cert_file", "/etc/oort/server.crt", "The TLS cert file")
	keyFile            = flag.String("key_file", "/etc/oort/server.key", "The TLS key file")
	port               = flag.Int("port", 8449, "The acctd server port")
	oortGroupHost      = flag.String("oortgrouphost", "127.0.0.1:6380", "host:port to use when connecting to oort group")
	insecureSkipVerify = flag.Bool("skipverify", true, "don't verify cert")
	superUserKey       = flag.String("superkey", "123456789", "Super User key used for authentication")
)

// FatalIf is just a lazy log/panic on error func
func FatalIf(err error, msg string) {
	if err != nil {
		grpclog.Fatalf("%s: %v", msg, err)
	}
}

func main() {
	flag.Parse()

	envtls := os.Getenv("OOHHC_ACCT_TLS")
	if envtls == "true" {
		*usetls = true
	}

	envoortghost := os.Getenv("OOHHC_OORT_GROUP_HOST")
	if envoortghost != "" {
		*oortGroupHost = envoortghost
	}

	envport := os.Getenv("OOHHC_ACCT_PORT")
	if envport != "" {
		p, err := strconv.Atoi(envport)
		if err != nil {
			log.Println("Did not send valid port from env:", err)
		} else {
			*port = p
		}
	}

	envcert := os.Getenv("OOHHC_ACCT_CERT_FILE")
	if envcert != "" {
		*certFile = envcert
	}

	envkey := os.Getenv("OOHHC_ACCT_KEY_FILE")
	if envkey != "" {
		*keyFile = envkey
	}

	envSuperKey := os.Getenv("OOHHC_ACCT_SUPERUSER_KEY")
	if envSuperKey != "" {
		*superUserKey = envSuperKey
	}

	envSkipVerify := os.Getenv("OOHHC_ACCT_SKIP_VERIFY")
	if envSkipVerify != "true" {
		*insecureSkipVerify = true
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	FatalIf(err, "Failed to bind to port")

	var opts []grpc.ServerOption
	if *usetls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		FatalIf(err, "Couldn't load cert from file")
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	s := grpc.NewServer(opts...)
	ws, err := NewAccountWS(*superUserKey, *oortGroupHost, *insecureSkipVerify)
	if err != nil {
		grpclog.Fatalln(err)
	}
	mb.RegisterAccountApiServer(s, NewAccountAPIServer(ws))
	grpclog.Printf("Starting up on %d...\n", *port)
	s.Serve(lis)
}
