package routes

import (
	"github.com/Mycunycu/gofermart/internal/handlers"
	customMiddleware "github.com/Mycunycu/gofermart/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	*chi.Mux
}

func NewRouter(h *handlers.Handler) *Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(customMiddleware.GzipCompress)
	r.Use(customMiddleware.GzipDecompress)

	r.Route("/api/user", func(r chi.Router) {
		r.Post("/register", h.Register())
		r.Post("/login", h.Login())
		r.Post("/orders", h.CreateOrder())
		r.Get("/orders", h.GetOrder())
		r.Get("/balance", h.GetBalance())
		r.Post("/balance/withdraw", h.WithdrawRequest())
		r.Post("/balance/withdrawals", h.WithdrawHistory())
	})

	return &Router{r}
}
