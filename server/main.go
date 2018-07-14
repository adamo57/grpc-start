package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/adamo57/grpc-start/user"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8081"
)

type server struct{}

// AddUser is eventually going to add a user to the DB
func (s *server) AddUser(ctx context.Context, in *pb.UserInfo) (*pb.AddUserMessage, error) {
	log.Printf("request recieved: %v", in)

	return &pb.AddUserMessage{Message: fmt.Sprintf("The Users info is: Name: %s, Age: %v, Email: %s", in.Name, in.Age, in.Email)}, nil
}

func (s *server) DelUser(ctx context.Context, in *pb.Username) (*pb.DelUserMessage, error) {
	return &pb.DelUserMessage{Message: "Not implemented"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", port)
	}

	log.Printf("INFO: server listening on %s", port)

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
