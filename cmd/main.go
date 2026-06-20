package main

import (
	"net/http"
	"strconv"

	"github.com/ch-load-balancer/internal/config"
	pkg "github.com/ch-load-balancer/pkg"
)

func main() {
	log := pkg.LoadLogger()
	cfg := config.LoadConfig(log)

	log.Info().Msg("Config loaded")

	log.Info().Msgf("Server listening on port: %d", cfg.Server.Port)
	if err := http.ListenAndServe(":"+strconv.Itoa(cfg.Server.Port), nil); err != nil {
		log.Fatal().Msgf("Error starting server: %v", err)
	}

}
