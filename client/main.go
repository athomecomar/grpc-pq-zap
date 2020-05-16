package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/athomecomar/grpc-pq-zap/pbauth"
	"google.golang.org/grpc"
)

const (
	address      = "localhost:50051"
	defaultEmail = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbauth.NewAuthClient(conn)

	// Contact the server and print out its response.
	email := defaultEmail
	if len(os.Args) > 1 {
		email = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SignIn(ctx, &pbauth.SignInRequest{Email: email, Password: "athome"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r.Users)
}
