package main
import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
	"net"
)

type greeter struct{
	helloworld.UnimplementedGreeterServer
}

func (greeter) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error,) {
	panic("not implemented")
}

func main(){
	var addr string
	flag.StringVar(&addr, "localhost:8080", "host:port of the server")
	log.Printf("Server (%s) starting up...", addr)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	creds := grpc.Creds(insecure.NewCredentials())
	server := grpc.NewServer(creds)
	helloworld.RegisterGreeterServer(server, greeter{})

	log.Println("Serving on ", listener.Addr())
	if err := server.Server(listener); err != nil {
		log.Fatal(err)
	}
}


