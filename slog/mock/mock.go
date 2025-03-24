package mock

import (
	"context"
	"log/slog"
)

func NewMockLogger() *slog.Logger {
	return slog.New(NewMock())
}

type Logger struct{}

func NewMock() *Logger {
	return &Logger{}
}

func (d Logger) Enabled(ctx context.Context, level slog.Level) bool {
	return false
}

func (d Logger) Handle(ctx context.Context, record slog.Record) error {
	return nil
}

func (d Logger) WithAttrs(attrs []slog.Attr) slog.Handler {
	return d
}

func (d Logger) WithGroup(name string) slog.Handler {
	return d
}
