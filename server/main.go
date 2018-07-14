// generate new proto file: protoc -I user/ user/userService.proto --go_out=plugins=grpc:user

package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	pb "github.com/adamo57/grpc-start/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8081"
)

type server struct{}

func (s *server) AddUser(ctx context.Context, in *pb.Info) (*pb.User, error) {
	return &pb.User{Name: in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", port)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
