package main

import (
	fmt "fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/devdinu/ideas"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func startGRPCServer() {
	stop := make(chan os.Signal, 2)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	listen, err := net.Listen("tcp", ":9988")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	ideas.RegisterCoolIdeasServer(server, &ideas.Service{})
	reflection.Register(server)
	fmt.Println("grpc server: listening on %s", "9988")
	go func() {
		if err := server.Serve(listen); err != nil {
			panic(err)
		}
	}()

	<-stop
	fmt.Println("stopping gRPC Server.")
	server.GracefulStop()
}

func main() {
	startGRPCServer()
}
