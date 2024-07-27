package authdto

type SignUpInputDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpOutputDto struct {
	UserId string `json:"userId"`
}
