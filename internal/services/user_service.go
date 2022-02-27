package services

import (
	"github.com/Mycunycu/gofermart/internal/repository"
)

var _ UserService = (*UserSvc)(nil)

type UserSvc struct {
	db repository.Repositorier
}

func NewUserService(db repository.Repositorier) *UserSvc {
	return &UserSvc{db: db}
}
