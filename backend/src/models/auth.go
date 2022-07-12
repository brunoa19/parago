package models

type AuthLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
