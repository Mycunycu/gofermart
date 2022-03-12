package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mycunycu/gofermart/internal/config"
	"github.com/Mycunycu/gofermart/internal/handlers"
	"github.com/Mycunycu/gofermart/internal/repository"
	"github.com/Mycunycu/gofermart/internal/routes"
	"github.com/Mycunycu/gofermart/internal/server"
	"github.com/Mycunycu/gofermart/internal/services"
)

func Run() error {
	cfg := config.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := repository.NewDatabase(ctx, cfg.DatabaseURI)
	if err != nil {
		return fmt.Errorf("error db connection: %v", err)
	}
	defer db.Close()

	err = db.Migrate(cfg.MigrationPath)
	if err != nil {
		return err
	}

	userSvc := services.NewUserService(db)

	handler := handlers.NewHandler(userSvc)
	router := routes.NewRouter(handler)
	srv := server.NewServer(ctx, cfg.ServerAddress, router)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)

	<-done

	timeoutCtx, timeoutCtxCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer func() {
		timeoutCtxCancel()
	}()

	if err := srv.Shutdown(timeoutCtx); err != nil {
		log.Fatalf("server shutdown failed:%+v", err)
	}

	if err == http.ErrServerClosed {
		err = nil
	}

	fmt.Println("Gracefull stopped")

	return nil
}
