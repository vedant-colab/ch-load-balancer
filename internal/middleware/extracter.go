package middleware

import (
	"context"
	"net/http"
)

type contextKey string

const UserIDKey contextKey = "userID"

func ExtractUserKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		userID := req.Header.Get("X-User-ID")
		if userID == "" {
			http.Error(w, "missing X-User-ID header", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(req.Context(), UserIDKey, userID)

		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
