package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Mycunycu/gofermart/internal/helpers"
	"github.com/Mycunycu/gofermart/internal/models"
	"github.com/Mycunycu/gofermart/internal/services"
	"github.com/go-chi/jwtauth/v5"
)

type Handler struct {
	userSvc   services.UserService
	tokenAuth *jwtauth.JWTAuth
}

func NewHandler(userSvc services.UserService, tokenAuth *jwtauth.JWTAuth) *Handler {
	return &Handler{userSvc: userSvc, tokenAuth: tokenAuth}
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var req models.RegisterRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.userSvc.Register(ctx, req)
		if err != nil {
			if errors.Is(err, helpers.ErrUnique) {
				http.Error(w, err.Error(), http.StatusConflict)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			return
		}

		_, tokenString, err := h.tokenAuth.Encode(models.UserClaims{"login": req.Login})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))
		w.WriteHeader(http.StatusOK)
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
