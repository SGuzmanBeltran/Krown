package types

import (
	"context"
	protouser "krown/services/genproto/user"

	"github.com/golang-jwt/jwt/v5"
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
	ValidateAuth(context.Context, *protouser.AuthRequest) (*protouser.AuthClaims ,error)
}

type JWTClaims struct {
	Username string
    Email string
	jwt.RegisteredClaims
}

func (c *JWTClaims) ParseToGRpcClaims() *protouser.AuthClaims {
	return &protouser.AuthClaims{
		Username: c.Username,
		Email: c.Email,
	}
}