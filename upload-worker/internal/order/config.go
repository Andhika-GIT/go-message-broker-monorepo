package order

import "github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"

func NewOrderModule(rmq *shared.RabbitMqConsumer) {
	StartWorker(rmq)
}
