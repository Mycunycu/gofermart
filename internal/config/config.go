package config

import (
	"flag"
	"log"
	"sync"

	"github.com/caarlos0/env"
)

type Config struct {
	ServerAddress  string `env:"RUN_ADDRESS" envDefault:":8080"`
	DatabaseURI    string `env:"DATABASE_URI" envDefault:"postgres://user:123@localhost:5432/gofermart?sslmode=disable"`
	AccrualAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	MigrationPath  string `env:"MIGRATION_PATH" envDefault:"file://internal/repository/migrations"`
}

var cfg Config

func New() Config {
	initialize()
	return cfg
}

func initialize() {
	var once sync.Once

	once.Do(func() {
		cfg = Config{}
		if err := env.Parse(&cfg); err != nil {
			log.Fatalf("initialize config error: %v", err)
		}

		flag.Func("a", "server address", func(value string) error {
			cfg.ServerAddress = value
			return nil
		})

		flag.Func("d", "database url", func(value string) error {
			cfg.DatabaseURI = value
			return nil
		})

		flag.Func("r", "accrual system address", func(value string) error {
			cfg.AccrualAddress = value
			return nil
		})

		flag.Parse()
	})
}
