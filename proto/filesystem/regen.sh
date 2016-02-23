#!/bin/bash 
set -x
# protoc --go_out=plugins=grpc:. *.proto
protoc --gofast_out=plugins=grpc:. *.proto
go install .
