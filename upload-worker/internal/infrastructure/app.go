package infrastructure

import (
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/user"
	"github.com/go-chi/chi/v5"
)

func InitApp() *chi.Mux {
	r := chi.NewRouter()

	v := NewViper()
	rmq, err := shared.NewRabbitMqConsumer(v)
	DB := NewDatabase(v)

	if err != nil {
		log.Fatalf("failed to initialize RabbitMQ connection: %v", err)
	}

	err = InitQueue(rmq)

	if err != nil {
		log.Fatalf("failed to bind RabbitMQ queues: %v", err)

	}

	user.NewUserModule(rmq, DB)

	return r
}
