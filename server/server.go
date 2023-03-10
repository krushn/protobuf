package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/krushn/protobuf/examples"

	"google.golang.org/grpc"
)

//proxy.golang.org/krushn/protobuf/examples

var (
	port = flag.Int("port", 50052, "The server port")
)

// server is used to implement Server
type server struct {
	pb.UnimplementedMessageGuideServer
}

// SayHello implements UnimplementedMessageGuideServer
func (s *server) SayHello(ctx context.Context, in *pb.Greet) (*pb.GreetReply, error) {
	log.Printf("Received: %v", in.GetName())

	var prename string

	if in.GetGender() == "male" {
		prename = "Mr."
	} else {
		prename = "Miss"
	}

	return &pb.GreetReply{Msg: "Hello, " + prename + " " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageGuideServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
