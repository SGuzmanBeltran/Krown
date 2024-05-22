package transport

import (
	"krown/services/tournament"
	"krown/services/tournament/handlers"
	"log"
	"net"

	"google.golang.org/grpc"
)

type TournamentgRPCServer struct {
	addr string
}

// NewTournamentGRPCServer creates a new gRPC server with the specified address.
func NewTournamentGRPCServer(addr string) *TournamentgRPCServer {
	return &TournamentgRPCServer{addr: addr}
}

func (s *TournamentgRPCServer) Run(tournamentStore *tournament.TournamentStore) error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	tournamentService := tournament.NewTournamentService(tournamentStore)
	handlers.NewGrpcTournamentService(grpcServer, tournamentService)

	log.Println("Starting Tournament gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
