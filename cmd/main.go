package main

import (
	"net/http"
	"strconv"

	"github.com/ch-load-balancer/internal/config"
	"github.com/ch-load-balancer/internal/handlers"
	"github.com/ch-load-balancer/internal/hashring"
	"github.com/ch-load-balancer/internal/middleware"
	"github.com/ch-load-balancer/internal/reverseproxy"
	pkg "github.com/ch-load-balancer/pkg"
)

func main() {
	log := pkg.LoadLogger()
	cfg := config.LoadConfig(log)
	log.Info().Msg("Config loaded")

	reverseproxy := reverseproxy.NewReverseConfig(cfg, log)
	err := reverseproxy.InitReverseProxies()
	if err != nil {
		log.Fatal().Msgf("Error initiating proxy: %v", err)
	}

	ring := hashring.NewHashRing(cfg, log)

	handler := handlers.Handler{
		Ring:         ring,
		ReverseProxy: reverseproxy,
	}

	extractKeyMiddleware := middleware.ExtractUserKey

	http.Handle("/debug", extractKeyMiddleware(http.HandlerFunc(handler.RoutingHandler)))
	http.Handle("/", extractKeyMiddleware(http.HandlerFunc(handler.ProxyHandler)))

	log.Info().Msgf("Server listening on port: %d", cfg.Server.Port)
	if err := http.ListenAndServe(":"+strconv.Itoa(cfg.Server.Port), nil); err != nil {
		log.Fatal().Msgf("Error starting server: %v", err)
	}

}
