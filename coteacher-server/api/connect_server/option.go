package connect_server

import (
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"

	"log/slog"
)

type option struct {
	logger      *slog.Logger
	entClient   *ent.Client
	FrontendURL string
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

func WithFrontendURL(url string) optionFunc {
	return func(o *option) {
		o.FrontendURL = url
	}
}
