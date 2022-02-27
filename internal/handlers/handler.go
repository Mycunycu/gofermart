package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Mycunycu/gofermart/internal/services"
)

type Handler struct {
	userSvc services.UserService
}

func NewHandler(userSvc services.UserService) *Handler {
	return &Handler{userSvc: userSvc}
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		w.Header().Set("content-type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(""))
	}
}

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		w.Header().Set("content-type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(""))
	}
}

func (h *Handler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		w.Header().Set("content-type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(""))
	}
}

func (h *Handler) GetOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		w.Header().Set("content-type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(""))
	}
}

func (h *Handler) GetBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		w.Header().Set("content-type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(""))
	}
}

func (h *Handler) WithdrawRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		w.Header().Set("content-type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(""))
	}
}

func (h *Handler) WithdrawHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		w.Header().Set("content-type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(""))
	}
}
