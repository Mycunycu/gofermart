package repository

import (
	"context"

	"github.com/Mycunycu/gofermart/internal/models"
)

type Repositorier interface {
	Save(context.Context, models.RegisterRequest) error
}
