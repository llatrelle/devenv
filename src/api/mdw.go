package api

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func IDCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		containerID := chi.URLParam(r, "id")
		ctx := context.WithValue(r.Context(), "id", containerID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
