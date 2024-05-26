package transport

import (
	"krown/common"
	"krown/services/scheduled"
	"krown/services/scheduled/handlers"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ScheduledgRPCServer struct {
	addr string
}

func NewScheduledGRPCServer(addr string) *ScheduledgRPCServer {
	return &ScheduledgRPCServer{addr: addr}
}

func (s *ScheduledgRPCServer) Run(scheduledStore *scheduled.ScheduledStore) error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	tournamentConn := common.NewGRPCClient(":9002")
	defer tournamentConn.Close()
	grpcServer := grpc.NewServer()
	scheduledService := scheduled.NewScheduledService(scheduledStore, tournamentConn)
	handlers.NewGrpcScheduledService(grpcServer, scheduledService)

	log.Println("Starting Scheduled gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}