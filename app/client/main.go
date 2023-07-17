package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"google.golang.org/grpc/"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/helloworld/helloworld"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "", "host:port of the server")
	flag.Parse()

	if addr == "" {
		addr = os.Getenv("GREETER_SERVER_ADDR")
	}
	if addr == "" {
		addr = "localhost:8123"
	}

	log.Println("Client is starting...", addr)

	ctx := context.Background()

	// TODO: Implementar comunicação segura com o uso do SPIRE
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())

	cliente, err := grpc.DialContext(ctx, creds)
	if err != nil {
		log.Fatal(err)
	}

	greeterClient := helloworld.NewGreeterClient(client)

	const interval = time.Second * 10
	log.Printf("We will be sending reuqests every %s s...\n", interval)
	// TODO: Implementar "send request"
	for {
	}
}
