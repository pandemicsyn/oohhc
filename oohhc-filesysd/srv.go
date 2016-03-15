package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"strings"
	"time"

	"golang.org/x/net/context"

	"github.com/gholt/brimtime"
	gp "github.com/pandemicsyn/oort/api/groupproto"
	"github.com/spaolacci/murmur3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// FileSystemWS the strurcture carrying all the extra stuff
type FileSystemWS struct {
	gaddr              string
	gopts              []grpc.DialOption
	gcreds             credentials.TransportAuthenticator
	insecureSkipVerify bool
	gconn              *grpc.ClientConn
	gc                 gp.GroupStoreClient
}

// NewFileSystemWS function used to create a new admin grpc web service
func NewFileSystemWS(gaddr string, insecureSkipVerify bool, grpcOpts ...grpc.DialOption) (*FileSystemWS, error) {
	// TODO: This all eventually needs to replaced with group rings
	var err error
	fs := &FileSystemWS{
		gaddr: gaddr,
		gopts: grpcOpts,
		gcreds: credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: insecureSkipVerify,
		}),
		insecureSkipVerify: insecureSkipVerify,
	}
	fs.gopts = append(fs.gopts, grpc.WithTransportCredentials(fs.gcreds))
	fs.gconn, err = grpc.Dial(fs.gaddr, fs.gopts...)
	if err != nil {
		return &FileSystemWS{}, err
	}
	fs.gc = gp.NewGroupStoreClient(fs.gconn)
	// TODO: this is copied from formicd so it doesn't reuse code.
	return fs, nil
}

// grpc Group Store functions
// getGroupClient ...
func (fsws *FileSystemWS) getGClient() {
	var opts []grpc.DialOption
	var creds credentials.TransportAuthenticator
	var err error
	creds = credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	opts = append(opts, grpc.WithTransportCredentials(creds))
	fsws.gconn, err = grpc.Dial(fsws.gaddr, opts...)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed to dial server: %s", err))
	}
	fsws.gc = gp.NewGroupStoreClient(fsws.gconn)
}

// lookupAccount ...
func (fsws *FileSystemWS) lookupGStore(g string) (string, error) {
	if fsws.gconn == nil {
		fsws.getGClient()
	}
	l := &gp.LookupGroupRequest{}
	l.KeyA, l.KeyB = murmur3.Sum128([]byte(g))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := fsws.gc.LookupGroup(ctx, l)
	if err != nil {
		return "", err
	}
	m := make([]string, len(res.Items))
	r := &gp.ReadRequest{}
	for k, v := range res.Items {
		r.KeyA = l.KeyA
		r.KeyB = l.KeyB
		r.ChildKeyA = v.ChildKeyA
		r.ChildKeyB = v.ChildKeyB
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		res, err := fsws.gc.Read(ctx, r)
		if err != nil {
			return "", err
		}
		m[k] = fmt.Sprintf("%s", res.Value)
	}
	return fmt.Sprintf(strings.Join(m, ",")), nil
}

// lookupAccount ...
func (fsws *FileSystemWS) writeGStore(g string, m string, p []byte) (string, error) {
	if fsws.gconn == nil {
		fsws.getGClient()
	}

	w := &gp.WriteRequest{}
	// prepare groupVal and memberVal
	w.KeyA, w.KeyB = murmur3.Sum128([]byte(g))
	w.ChildKeyA, w.ChildKeyB = murmur3.Sum128([]byte(m))
	w.Value = p
	w.TimestampMicro = brimtime.TimeToUnixMicro(time.Now())
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := fsws.gc.Write(ctx, w)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("WRITE TSM: %d\nTSM: %d", w.TimestampMicro, res.TimestampMicro), nil
}

// lookupAccount ...
func (fsws *FileSystemWS) getGStore(g string, m string) (string, error) {
	if fsws.gconn == nil {
		fsws.getGClient()
	}
	// TODO:
	r := &gp.ReadRequest{}
	r.KeyA, r.KeyB = murmur3.Sum128([]byte(g))
	r.ChildKeyA, r.ChildKeyB = murmur3.Sum128([]byte(m))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := fsws.gc.Read(ctx, r)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", res.Value), nil
}
