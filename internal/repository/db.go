package repository

import (
	"context"
	"errors"

	"github.com/Mycunycu/gofermart/internal/models"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ Repositorier = (*Database)(nil)

type Database struct {
	*pgxpool.Pool
}

func NewDatabase(ctx context.Context, connStr string) (*Database, error) {

	pool, err := connectDB(ctx, connStr)
	if err != nil {
		return nil, errors.New("db connection error")
	}

	return &Database{pool}, nil
}

func connectDB(ctx context.Context, connStr string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func (d *Database) Migrate(source string) error {
	m, err := migrate.New(source, d.Config().ConnString())
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	}
	return nil
}

func (d *Database) Save(ctx context.Context, e models.RegisterRequest) error {
	sql := "INSERT INTO customer VALUES (default, $1, $2)"
	_, err := d.Exec(ctx, sql, e.Login, e.Password)
	if err != nil {
		return err
	}
	return nil
}
