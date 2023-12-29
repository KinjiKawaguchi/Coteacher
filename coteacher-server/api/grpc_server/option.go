package grpc_server

import (
	"coteacher/domain/repository/ent"

	"golang.org/x/exp/slog"
)

type option struct {
	logger    *slog.Logger
	entClient *ent.Client
}

func defaultOption() *option {
	return &option{
		logger: slog.Default(),
	}
}

type optionFunc func(*option)

func WithLogger(logger *slog.Logger) optionFunc {
	return func(o *option) {
		o.logger = logger
	}
}

func WithEntClient(c *ent.Client) optionFunc {
	return func(o *option) {
		o.entClient = c
	}
}
