package main

import (
	"context"
	pb "github.com/root27-dev/grpc-k8s/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedAddServiceServer
}

func main() {

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {

		log.Fatalf("Failed to listen: %v", err)

	}

	srv := grpc.NewServer()

	pb.RegisterAddServiceServer(srv, &server{})

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{Result: in.A + in.B}, nil
}
