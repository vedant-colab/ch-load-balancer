package pkg

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LoadLogger() *zerolog.Logger {
	return &log.Logger
}
