package main

import (
	"context"
	"log"
	"net"

	"github.com/athomecomar/grpc-pq-zap/authconf"
	"github.com/athomecomar/grpc-pq-zap/pb/pbauth"
	"google.golang.org/grpc"
)

type server struct {
	pbauth.UnimplementedAuthServer
}

func (s *server) SignIn(ctx context.Context, in *pbauth.SignInRequest) (*pbauth.SignInResponse, error) {
	log.Printf("Received: %v", in.GetEmail())
	return &pbauth.SignInResponse{Users: []*pbauth.SignInUser{
		{Token: "bar", Role: "foo"},
		{Token: "quux", Role: "baz"},
	}}, nil
}

func main() {
	port := authconf.GetPORT()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbauth.RegisterAuthServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
