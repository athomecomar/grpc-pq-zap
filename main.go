package main

import (
	"context"
	"log"
	"net"

	"github.com/athomecomar/grpc-pq-zap/pb/pbuser"
	"github.com/athomecomar/grpc-pq-zap/userconf"
	"google.golang.org/grpc"
)

type server struct {
	pbuser.UnimplementedUserServer
}

func (s *server) SignIn(ctx context.Context, in *pbuser.SignInRequest) (*pbuser.SignInResponse, error) {
	log.Printf("Received: %v", in.GetEmail())
	return &pbuser.SignInResponse{Users: []*pbuser.SignInUser{
		{Token: "bar", Role: "foo"},
		{Token: "quux", Role: "baz"},
	}}, nil
}

func main() {
	port := userconf.GetPORT()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbuser.RegisterUserServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
