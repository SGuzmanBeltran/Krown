package types

import (
	"context"
	protouser "krown/services/genproto/user"
)

type LoginUserPayload struct {
	Email string  `json:"email"`
	Password string `json:"password"`
}

type ServiceResponse struct {
	Message string `json:"message"`
	Status int `json:"status"`
}


type Response struct {
	Message string `json:"message"`
}

type ResponseToken struct {
	Token string `json:"token"`
}

type UserService interface {
	ValidateAuth(context.Context, *protouser.AuthRequest) error
}