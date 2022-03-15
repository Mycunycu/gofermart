package models

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserClaims map[string]interface{}
