package logger

import (
	"io"
	"os"
	"time"

	"github.com/lingojack/proj_template/config"
	"github.com/rs/zerolog"
)

func New(cfg *config.Config) zerolog.Logger {
	level, err := zerolog.ParseLevel(cfg.Log.Level)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	var w io.Writer
	if cfg.Log.Format == "text" {
		w = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	} else {
		w = os.Stdout
	}

	if cfg.Log.Output == "file" && cfg.Log.FilePath != "" {
		f, err := os.OpenFile(cfg.Log.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			w = f
		}
	}

	return zerolog.New(w).With().Timestamp().Logger()
}
