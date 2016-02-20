package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"

	"github.com/gholt/brimtime"
	gp "github.com/pandemicsyn/oort/api/groupproto"
	"github.com/spaolacci/murmur3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// AccountWS the strurcture carrying all the extra stuff
type AccountWS struct {
	superKey           string
	gaddr              string
	gopts              []grpc.DialOption
	gcreds             credentials.TransportAuthenticator
	insecureSkipVerify bool
	gconn              *grpc.ClientConn
	gc                 gp.GroupStoreClient
}

// NewAccountWS function used to create a new admin grpc web service
func NewAccountWS(superkey string, gaddr string, insecureSkipVerify bool, grpcOpts ...grpc.DialOption) (*AccountWS, error) {
	// TODO: This all eventually needs to replaced with group rings
	var err error
	a := &AccountWS{
		superKey: superkey,
		gaddr:    gaddr,
		gopts:    grpcOpts,
		gcreds: credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: insecureSkipVerify,
		}),
		insecureSkipVerify: insecureSkipVerify,
	}
	a.gopts = append(a.gopts, grpc.WithTransportCredentials(a.gcreds))
	a.gconn, err = grpc.Dial(a.gaddr, a.gopts...)
	if err != nil {
		return &AccountWS{}, err
	}
	a.gc = gp.NewGroupStoreClient(a.gconn)
	// TODO: this is copied from formicd so it doesn't reuse code.
	return a, nil
}

// grpc Group Store functions
// getGroupClient ...
func (aws *AccountWS) getGClient() {
	var opts []grpc.DialOption
	var creds credentials.TransportAuthenticator
	var err error
	creds = credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	opts = append(opts, grpc.WithTransportCredentials(creds))
	aws.gconn, err = grpc.Dial(aws.gaddr, opts...)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed to dial server: %s", err))
	}
	aws.gc = gp.NewGroupStoreClient(aws.gconn)
}

// lookupAccount ...
func (aws *AccountWS) lookupGStore(g string, m string) (string, error) {
	if aws.gconn == nil {
		aws.getGClient()
	}
	// TODO:
	return "ok", nil
}

// lookupAccount ...
func (aws *AccountWS) writeGStore(g string, m string, p []byte) (string, error) {
	if aws.gconn == nil {
		aws.getGClient()
	}

	w := &gp.WriteRequest{}
	// prepare groupVal and memberVal
	w.KeyA, w.KeyB = murmur3.Sum128([]byte(g))
	w.NameKeyA, w.NameKeyB = murmur3.Sum128([]byte(m))
	w.Value = p
	w.Tsm = brimtime.TimeToUnixMicro(time.Now())
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := aws.gc.Write(ctx, w)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("WRITE TSM: %d\nTSM: %d", w.Tsm, res.Tsm), nil
}

// lookupAccount ...
func (aws *AccountWS) getGStore(g string, m string) (string, error) {
	if aws.gconn == nil {
		aws.getGClient()
	}
	// TODO:
	r := &gp.ReadRequest{}
	r.KeyA, r.KeyB = murmur3.Sum128([]byte(g))
	r.NameKeyA, r.NameKeyB = murmur3.Sum128([]byte(m))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := aws.gc.Read(ctx, r)
	if err != nil {
		return "", err
	}
	fmt.Println("RETURNED VALUE FROM THE GROUP STORE\n", res.Value)
	return fmt.Sprintf("%s", res.Value), nil
}
