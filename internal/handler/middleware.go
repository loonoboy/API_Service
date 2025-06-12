package handler

import (
	resp "API_Service/internal/response"
	"context"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

type contextKey string

const (
	authorizationHeader            = "Authorization"
	userCtxKey          contextKey = "userId"
)

func (h *Handler) userIdentity() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get(authorizationHeader)
			if header == "" {
				w.WriteHeader(http.StatusUnauthorized)
				render.JSON(w, r, resp.Error("empty authorization header"))
				return
			}
			headerParts := strings.Split(header, " ")
			if len(headerParts) != 2 {
				w.WriteHeader(http.StatusUnauthorized)
				render.JSON(w, r, resp.Error("invalid authorization header"))
				return
			}

			userId, err := h.service.Authorization.ParseToken(headerParts[1])
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				render.JSON(w, r, resp.Error("Parse Token failed"))
				return
			}
			ctx := context.WithValue(r.Context(), userCtxKey, userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
