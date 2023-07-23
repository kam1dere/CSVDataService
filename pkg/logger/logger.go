package logger

import (
	"os"

	"github.com/kam1dere/CSVDataService/config"
	"github.com/rs/zerolog"
)

func NewLogger(cfg *config.Config) zerolog.Logger {
	return zerolog.New(os.Stdout).With().Timestamp().Str("ServiceName", cfg.Logger.ServiceName).Caller().Logger()
}
