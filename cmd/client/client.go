package main

import (
	"context"
	"fmt"
	"log"

	"github.com/devdinu/ideas"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9988", grpc.WithInsecure())
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
