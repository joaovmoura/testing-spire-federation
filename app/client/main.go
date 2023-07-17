package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/peer"
)

func sendRequest(ctx context.Context, client helloworld.GreeterClient) {
	peer := new(peer.Peer)
	res, err := client.SayHello(ctx, &helloworld.HelloRequest{}, grpc.Peer(peer))
	if err != nil {
		log.Printf("Failed to send hello request. Error: %v\n", err)
		return
	}

	// TODO: Implementar identidade do server com SVID's
	serverId := "UnknownServer"

	log.Printf("%s said %q", serverId, res.Message)
}
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

	client, err := grpc.DialContext(ctx, addr, creds)
	if err != nil {
		log.Fatal(err)
	}

	greeterClient := helloworld.NewGreeterClient(client)

	const interval = time.Second * 10
	log.Printf("We will be sending reuqests every %s s...\n", interval)
	// TODO: Implementar "send request"
	for {
		sendRequest(ctx, greeterClient)
		time.Sleep(interval)
	}
}
