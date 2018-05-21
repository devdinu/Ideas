package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/devdinu/ideas"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 9988, "port for server to listen")
	host := flag.String("host", "localhost", "host for server to listen")
	flag.Parse()

	makeCall(*host, *port)
}

func makeCall(host string, port int) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		fmt.Println("conn problem", err)
	}
	cli := ideas.NewCoolIdeasClient(conn)
	u := &ideas.User{Id: 1235}
	res, err := cli.GetIdeas(context.Background(), u)
	if err != nil {
		log.Fatalf("could not get count: %v", err)
		return
	}
	log.Printf("Response From Server: %v", res.GetIdeas())
}
