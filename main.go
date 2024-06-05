package main

import (
	"flag"
	"log"
	"net"
	"strings"
	"time"

	pb "./se_coordinator/"

	"google.golang.org/grpc"
)

func setupClients(peers []string, peerList []string) map[string]pb.SeCoordinatorClient {
	connections := make(map[string]pb.SeCoordinatorClient)

	for i, peer := range peers {
		var conn *grpc.ClientConn
		var err error
		for retry := 0; retry < 5; retry++ { // Retry 5 times
			conn, err = grpc.Dial(peerList[i], grpc.WithInsecure())
			if err == nil {
				break
			}
			log.Printf("Failed to connect to %s (%s): %v. Retrying...", peer, peerList[i], err)
			time.Sleep(2 * time.Second) // Wait for 2 seconds before retrying
		}
		if err != nil {
			log.Printf("Could not connect to %s (%s) after multiple attempts: %v", peer, peerList[i], err)
			continue // Skip this peer if still failing
		}
		connections[peer] = pb.NewSeCoordinatorClient(conn)
	}

	return connections
}

func main() {
	var (
		name        string
		port        string
		peersStr    string
		peerListStr string
	)

	flag.StringVar(&name, "name", "", "current server/datacenter name")
	flag.StringVar(&port, "port", "50051", "port to run the server on")
	flag.StringVar(&peersStr, "peers", "", "comma-separated list of peer names")
	flag.StringVar(&peerListStr, "peerList", "", "comma-separated list of peer addresses")
	flag.Parse()

	peers := strings.Split(peersStr, ",")
	peerList := strings.Split(peerListStr, ",")

	if len(peers) != len(peerList) {
		log.Fatalf("number of peers and peer addresses must match")
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}
	s := grpc.NewServer()

	tpLayer := NewTpLayer()
	seCoordinator := NewSeCoordinator(tpLayer)

	pb.RegisterSeCoordinatorServer(s, &SeCoordinatorServerImpl{
		DatacenterID: name,
		Clients:      setupClients(peers, peerList),
		Coordinator:  seCoordinator,
	})

	log.Printf("Starting server %s on port %s", name, port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
