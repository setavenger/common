package zlog

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	// Shorthand logger
	L           zerolog.Logger
	Environment string = "development"
)

func init() {

	if val := os.Getenv("ENVIRONMENT"); val != "" {
		Environment = val
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixNano

	if Environment != "production" {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		L = zerolog.New(output).With().Timestamp().Logger()

	}
}
