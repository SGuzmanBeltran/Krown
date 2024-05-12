package types

type LoginUserPayload struct {
	Email string  `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}

type ResponseToken struct {
	Token string `json:"token"`
}