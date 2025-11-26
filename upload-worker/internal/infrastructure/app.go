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
	DB := NewDatabase(&cfg.Database)

	sftp, err := NewSFTPClient(&cfg.SftpClient)

	if err != nil {
		log.Fatalf("failed to initialize sftp: %v", err)
	}

	rmq, err := shared.NewRabbitMqConsumer(cfg.RabbitMQConnectURL)

	if err != nil {
		log.Fatalf("failed to initialize RabbitMQ connection: %v", err)
	}

	err = InitQueue(rmq, cfg)

	if err != nil {
		log.Fatalf("failed to bind RabbitMQ queues: %v", err)

	}

	userModule := user.NewUserModule(rmq, DB, &cfg.RabbitMQQueue, sftp)
	order.NewOrderModule(rmq, DB, userModule.UserUseCase, &cfg.RabbitMQQueue, sftp)

	return r
}
