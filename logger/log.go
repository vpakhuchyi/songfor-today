package logger

import (
	"os"

	"golang.org/x/exp/slog"
)

func SetLevel(level string) {
	var l slog.Level
	switch level {
	case "INFO":
		l = slog.LevelInfo
	case "DEBUG":
		l = slog.LevelDebug
	}

	slog.SetDefault(slog.New(slog.HandlerOptions{Level: l}.NewTextHandler(os.Stderr)))
}
