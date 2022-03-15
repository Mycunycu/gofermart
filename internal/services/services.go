package services

import (
	"context"

	"github.com/Mycunycu/gofermart/internal/models"
)

type UserService interface {
	Register(context.Context, models.RegisterRequest) error
}
