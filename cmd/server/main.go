package main

import (
	"flag"
	fmt "fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/devdinu/ideas"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func startGRPCServer(port int) {
	stop := make(chan os.Signal, 2)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	ideas.RegisterCoolIdeasServer(server, &ideas.Service{})
	reflection.Register(server)
	fmt.Printf("grpc server: listening on %s\n", "9988")
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
	port := flag.Int("port", 9988, "port for server to listen")
	flag.Parse()

	startGRPCServer(*port)
}
