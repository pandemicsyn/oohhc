package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	mb "github.com/letterj/oohhc/proto/account"
	gp "github.com/pandemicsyn/oort/api/groupproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"net"
)

var (
	usetls             = flag.Bool("tls", true, "Connection uses TLS if true, else plain TCP")
	certFile           = flag.String("cert_file", "/etc/oort/server.crt", "The TLS cert file")
	keyFile            = flag.String("key_file", "/etc/oort/server.key", "The TLS key file")
	port               = flag.Int("port", 8449, "The server port")
	oortGroupHost      = flag.String("oortgrouphost", "127.0.0.1:6380", "host:port to use when connecting to oort group")
	insecureSkipVerify = flag.Bool("skipverify", true, "don't verify cert")
	superUserKey       = flag.String("superkey", "123456789", "Super User key used for authentication")
)

// AccountWS the strurcture carrying all the extra stuff
type AccountWS struct {
	superKey           string
	gaddr              string
	gopts              []grpc.DialOption
	gcreds             credentials.TransportAuthenticator
	insecureSkipVerify bool
	gconn              *grpc.ClientConn
	gclient            gp.GroupStoreClient
}

// NewAccountWS function used to create a new admin grpc web service
func NewAccountWS(superkey string, gaddr string, insecureSkipVerify bool, grpcOpts ...grpc.DialOption) (*AccountWS, error) {
	// TODO: This all eventually needs to replaced with group rings
	var err error
	o := &AccountWS{
		superKey: superkey,
		gaddr:    gaddr,
		gopts:    grpcOpts,
		gcreds: credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: insecureSkipVerify,
		}),
		insecureSkipVerify: insecureSkipVerify,
	}
	o.gopts = append(o.gopts, grpc.WithTransportCredentials(o.gcreds))
	o.gconn, err = grpc.Dial(o.gaddr, o.gopts...)
	if err != nil {
		return &AccountWS{}, err
	}
	o.gclient = gp.NewGroupStoreClient(o.gconn)
	// TODO: this is copied from formicd so it doesn't reuse code.
	return o, nil
}

// FatalIf is just a lazy log/panic on error func
func FatalIf(err error, msg string) {
	if err != nil {
		grpclog.Fatalf("%s: %v", msg, err)
	}
}

func main() {
	flag.Parse()

	envtls := os.Getenv("OOHHC_ACCOUNT_TLS")
	if envtls == "true" {
		*usetls = true
	}

	envoortghost := os.Getenv("OOHHC_OORT_GROUP_HOST")
	if envoortghost != "" {
		*oortGroupHost = envoortghost
	}

	envport := os.Getenv("OOHHC_PORT")
	if envport != "" {
		p, err := strconv.Atoi(envport)
		if err != nil {
			log.Println("Did not send valid port from env:", err)
		} else {
			*port = p
		}
	}

	envcert := os.Getenv("OOHHC_CERT_FILE")
	if envcert != "" {
		*certFile = envcert
	}

	envkey := os.Getenv("OOHHC_KEY_FILE")
	if envkey != "" {
		*keyFile = envkey
	}

	envSuperKey := os.Getenv("OOHHC_SUPERUSER_KEY")
	if envSuperKey != "" {
		*superUserKey = envSuperKey
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
