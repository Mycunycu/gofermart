package services

import (
	"context"

	"github.com/Mycunycu/gofermart/internal/models"
	"github.com/Mycunycu/gofermart/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var _ UserService = (*UserSvc)(nil)

type UserSvc struct {
	db repository.Repositorier
}

func NewUserService(db repository.Repositorier) *UserSvc {
	return &UserSvc{db: db}
}

func (u *UserSvc) Register(ctx context.Context, user models.RegisterRequest) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}

	user.Password = string(hashed)

	err = u.db.Save(ctx, user)
	if err != nil {
		//
	}

	return nil
}
