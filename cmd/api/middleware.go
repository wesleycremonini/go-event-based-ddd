package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/wesleycremonini/go-event-based-ddd/internal/domain"
	"github.com/wesleycremonini/go-event-based-ddd/internal/response"
	"go.uber.org/zap"
)

type Middlewares struct {
	users domain.UserRepository
}

type contextKey string

// recoverPanic handle panics for the application not to finish.
func (m Middlewares) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				zap.L().Error("recoverPanic", zap.Error(fmt.Errorf("%s", err)))
				response.Error(w, http.StatusInternalServerError, "internal error", nil)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// loggedUser requires x-logged-user header to be present and valid.
func (m Middlewares) loggedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		loggedUser, err := uuid.Parse(r.Header.Get("x-logged-user"))
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "invalid x-logged-user header", nil)
			return
		}

		user, err := m.users.GetByUUID(r.Context(), loggedUser)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "invalid x-logged-user header", nil)
			return
		}

		const userKey contextKey = "user"
		c := context.WithValue(r.Context(), userKey, user)
		r = r.WithContext(c)

		next.ServeHTTP(w, r)
	})
}

func (m Middlewares) allowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,x-customer-token,x-logged-user")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
