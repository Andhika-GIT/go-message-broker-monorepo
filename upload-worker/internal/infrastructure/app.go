package infrastructure

import (
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/order"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/user"
	"github.com/go-chi/chi/v5"
)

func InitApp() *chi.Mux {
	r := chi.NewRouter()

	v := NewViper()
	cfg := shared.InitConfig(v)
	rmq, err := shared.NewRabbitMqConsumer(cfg.RabbitMQConnectURL)
	DB := NewDatabase(&cfg.Database)

	if err != nil {
		log.Fatalf("failed to initialize RabbitMQ connection: %v", err)
	}

	err = InitQueue(rmq, cfg)

	if err != nil {
		log.Fatalf("failed to bind RabbitMQ queues: %v", err)

	}

	userModule := user.NewUserModule(rmq, DB)
	order.NewOrderModule(rmq, DB, userModule.UserUseCase, &cfg.RabbitMQQueue)

	return r
}
