package user

import "github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"

func NewUserModule(rmq *shared.RabbitMqConsumer) {
	StartWorker(rmq)
}
