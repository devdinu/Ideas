protoc --go_out=$GOPATH/src ideas.proto
protoc -I . ./ideas.proto --go_out=plugins=grpc:$GOSRC
