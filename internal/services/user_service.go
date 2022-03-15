package services

import (
	"context"
	"errors"

	"github.com/Mycunycu/gofermart/internal/helpers"
	"github.com/Mycunycu/gofermart/internal/models"
	"github.com/Mycunycu/gofermart/internal/repository"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
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
	var targetErr *pgconn.PgError
	if errors.As(err, &targetErr) && targetErr.Code == pgerrcode.UniqueViolation {
		return helpers.ErrUnique
	}

	return nil
}
