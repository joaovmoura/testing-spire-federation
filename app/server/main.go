package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/examples/helloworld/helloworld"
)

type greeter struct {
	helloworld.UnimplementedGreeterServer
}

func (greeter) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	clientId := "UnknownClientId" // Não é possível saber o id do cliente ainda. Isso vai mudar com o uso do spire.

	log.Printf("%s has requested that I say hello world to %q...", clientId, req.Name)

	return &helloworld.HelloReply{
		Message: fmt.Sprintf("On behalf of %s, hello %s.", clientId, req.Name),
	}, nil
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "localhost:8123", "host:port of the server")
	log.Printf("Server (%s) starting up...", addr)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	//TODO: implementar comunicaçã segura usando SVID
	creds := grpc.Creds(insecure.NewCredentials())
	server := grpc.NewServer(creds)
	helloworld.RegisterGreeterServer(server, greeter{})

	log.Println("Serving on ", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
