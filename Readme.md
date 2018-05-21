### Ideas - gRPC Service

A basic CRUD service, to explore and explain gRPC.

`ideas.proto` contains the definition of service

```
service CoolIdeas {
    rpc SubmitIdea(Idea) returns (IdeaResponse) {}
    rpc GetIdeas(User) returns (Ideas) {}
}
```

Service struct in `service.go` satisifies the above interface and can be registerd and run as gRPC Service.
```
	ideas.RegisterCoolIdeasServer(server, &ideas.Service{}) // file: server/main.go
```
Client makes request to the server, by creating a client and calling `GetIdeas` or `SubmitIdea`

```
    ideas.NewCoolIdeasClient(conn).GetIdeas(ctx, user) // file: client/client.go
```

### Protobuf
Instead of using JSON to send payload to request, gRPC uses protocol buffers (protobuf), its smaller and performant.
```
message Idea {
    User user_id = 1;
    string title = 2;
    string description = 3;
}
```

### Requirements
- You should 've go installed, check with `go version`
- `protoc --version` to verify protoc installation, along with `protoc-gen-go`
or install it by

```
brew install protobuf
go get -u github.com/golang/protobuf/protoc-gen-go
```

### Running the Service
* `cmd/server` dir has server code, `cmd/client` has the client code.
* Generate the ideas.pb.proto which contains the service definition and protobuf, by running

```
protoc -I . ./ideas.proto --go_out=plugins=grpc:$GOSRC
```
* `go build && ./server` in server dir to run server
* `go build && ./client` in client dir to run client

you can optionally pass the port flag while running server/client to listen/dial on that port

### Resources
* [protobuf vs http benchmark](http://pliutau.com/benchmark-grpc-protobuf-vs-http-json/)
* [why use protobuf](https://developers.google.com/protocol-buffers/docs/gotutorial#why-use-protocol-buffers)
* [compile protobuf](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers)
