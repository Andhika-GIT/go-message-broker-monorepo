package infrastructure

import (
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/order"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/user"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/worker"
	"github.com/go-chi/chi/v5"
)

func InitApp() *chi.Mux {
	r := NewRouter()

	v := NewViper()
	cfg := shared.InitConfig(v)
	db := NewDatabase(&cfg.Database)
	rmq, err := shared.NewRabbitMqProducer(cfg.RabbitMQConnectURL)

	if err != nil {
		log.Fatalf("failed to initialize RabbitMQ connection: %v", err)
	}

	err = InitQueue(rmq, cfg)

	if err != nil {
		log.Fatalf("failed to bind RabbitMQ queues: %v", err)

	}

	sftpClient, err := NewSFTPClient(&cfg.SftpClient)

	if err != nil {
		log.Fatalf("failed to bind to sftp client %v", err)
	}

	uploadWorker := worker.NewUploadWorker(sftpClient, rmq, 3)
	order.NewOrderModule(r, rmq, uploadWorker, db)
	user.NewUserModule(r, rmq, uploadWorker, db)

	go uploadWorker.Start()

	return r
}
