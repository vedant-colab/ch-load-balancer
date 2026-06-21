package handlers

import (
	"log"
	"net/http"

	"github.com/ch-load-balancer/internal/middleware"
)

func RoutingHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "user id not found", http.StatusBadRequest)
		return
	}

	log.Println(userId)
	// will complete later
}
