package transport

import (
	"krown/services/user"
	handler "krown/services/user/handlers"
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

// NewGRPCServer creates a new gRPC server with the specified address.
func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run(userStore *user.Store) error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	userService := user.NewUserService(userStore)
	handler.NewGrpcUserService(grpcServer, userService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
