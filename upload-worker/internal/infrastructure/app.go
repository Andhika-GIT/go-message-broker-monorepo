package infrastructure

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/go-chi/chi/v5"
)

func InitApp() *chi.Mux {
	r := chi.NewRouter()

	v := NewViper()
	shared.NewRabbitMqProducer(v)

	return r
}
