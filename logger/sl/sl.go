package sl

import (
	"errors"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var ErrorWrongEnv = errors.New("select one of available envs: local, dev, prod")

func SetupLogger(env string) (*slog.Logger, error) {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	if logger == nil {
		return nil, ErrorWrongEnv
	}

	logger = logger.With(slog.String("env", env))
	return logger, nil
}

func Err(err error) slog.Attr {
	return ErrMsg(err.Error())
}

func ErrMsg(error string) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(error),
	}
}
