package infrastructure

import (
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

func InitQueue(rmq *shared.RabbitMqConsumer) error {
	log.Println("Initializing RabbitMQ...")

	// Declare exchange
	if err := rmq.DeclareExchange(shared.ExchangeGoApp, "direct"); err != nil {
		return err
	}

	// --- User Queues ---
	if err := rmq.QueueBind(shared.QueueUserDirectImport, shared.ExchangeGoApp, shared.RoutingKeyUserDirectImport); err != nil {
		return err
	}

	if err := rmq.QueueBind(shared.QueueUserSftpImport, shared.ExchangeGoApp, shared.RoutingKeyUserSftpImport); err != nil {
		return err
	}

	return nil
}
