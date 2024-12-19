package slog_discard

import (
	"context"
	"log/slog"
)

func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

type DiscardHandler struct{}

func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

func (d DiscardHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return false
}

func (d DiscardHandler) Handle(ctx context.Context, record slog.Record) error {
	return nil
}

func (d DiscardHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return d
}

func (d DiscardHandler) WithGroup(name string) slog.Handler {
	return d
}
