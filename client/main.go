package main

import (
	"context"
	"log"
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

	user := &pb.UserInfo{
		Name:  "Adam Ouellette",
		Age:   24,
		Email: "adamouellette57@gmail.com",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ar, err := c.AddUser(ctx, user)
	if err != nil {
		log.Fatalf("could not add user: %v", err)
	}

	log.Printf("Reponse: %s", ar.Message)

	dr, err := c.DelUser(ctx, &pb.Username{Username: "adamo57"})
	if err != nil {
		log.Fatalf("could not delete user: %v", err)
	}

	log.Printf("Repsonse: %s", dr.Message)
}
