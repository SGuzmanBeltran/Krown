package handler

import (
	"context"
	"krown/common/types"
	protouser "krown/services/genproto/user"

	"google.golang.org/grpc"
)

type UserGrpcHanlder struct {
	userService types.UserService
	protouser.UnimplementedAuthServiceServer
}

func NewGrpcUserService(grpc *grpc.Server, userService types.UserService) {
	gRpcHandler := &UserGrpcHanlder{userService:userService}

	protouser.RegisterAuthServiceServer(grpc, gRpcHandler)
}

func (h *UserGrpcHanlder) ValidateAuth(ctx context.Context, req *protouser.AuthRequest) (*protouser.AuthResponse, error) {
	h.userService.ValidateAuth(ctx, req)

	err := h.userService.ValidateAuth(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &protouser.AuthResponse{
		Valid: true,
	}

	return res, nil
}