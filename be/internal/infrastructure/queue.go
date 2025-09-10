package infrastructure

import (
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

func InitQueue(rmq *shared.RabbitMqProducer) error {
	log.Println("Initializing RabbitMQ...")

	// Declare exchange
	if err := rmq.DeclareExchange(shared.ExchangeGoApp, "direct"); err != nil {
		return err
	}

	// --- User Queues ---
	if err := rmq.QueueBind(shared.QueueUserImport, shared.ExchangeGoApp, shared.RoutingKeyUserImport); err != nil {
		return err
	}

	if err := rmq.QueueBind(shared.QueueUserExport, shared.ExchangeGoApp, shared.RoutingKeyUserExport); err != nil {
		return err
	}

	// --- Order Queues ---
	if err := rmq.QueueBind(shared.QueueOrderImport, shared.ExchangeGoApp, shared.RoutingKeyOrderImport); err != nil {
		return err
	}

	if err := rmq.QueueBind(shared.QueueOrderExport, shared.ExchangeGoApp, shared.RoutingKeyOrderExport); err != nil {
		return err
	}

	return nil
}
