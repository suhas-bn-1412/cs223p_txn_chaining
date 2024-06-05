package main

import (
	"log"
	"net"
	"os"
	pb "path/to/proto"

	"google.golang.org/grpc"
)

func setupClients() map[string]pb.SeCoordinatorClient {
	connections := make(map[string]pb.SeCoordinatorClient)

	connB, err := grpc.Dial("datacenter_b:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	connections["B"] = pb.NewSeCoordinatorClient(connB)

	connC, err := grpc.Dial("datacenter_c:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	connections["C"] = pb.NewSeCoordinatorClient(connC)

	connD, err := grpc.Dial("datacenter_d:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	connections["D"] = pb.NewSeCoordinatorClient(connD)

	return connections
}

func main() {
	port := os.Getenv("PORT")
	datacenterID := os.Getenv("DATACENTER_ID")

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSeCoordinatorServer(s, &SeCoordinatorServerImpl{
		DatacenterID: datacenterID,
		Clients:      setupClients(),
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
