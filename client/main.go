package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/adamo57/grpc-start/user"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:8081"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	name := "Adam Ouellette"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddUser(ctx, &pb.Info{Name: name})
	if err != nil {
		log.Fatalf("could not add user: %v", err)
	}

	log.Printf("User: %s", r.Name)
}
