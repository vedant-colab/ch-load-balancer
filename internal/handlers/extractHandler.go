package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ch-load-balancer/internal/hashring"
	"github.com/ch-load-balancer/internal/middleware"
	"github.com/ch-load-balancer/internal/reverseproxy"
)

type Handler struct {
	Ring         *hashring.HashRing
	ReverseProxy *reverseproxy.ReverseProxyConfig
}

func (h Handler) RoutingHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "user id not found", http.StatusBadRequest)
		return
	}

	log.Println(userId)
	svr := h.Ring.Lookup(userId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]any{"message": "success", "server": svr}
	json.NewEncoder(w).Encode(response)

}

func (h Handler) ProxyHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "user id not found", http.StatusBadRequest)
		return
	}

	log.Println(userId)
	svr := h.Ring.Lookup(userId)
	url := "http://" + svr.Host + ":" + strconv.Itoa(svr.Port)
	h.ReverseProxy.ForwardRequest(url, w, r)

}
