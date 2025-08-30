package order

import (
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

func StartWorker(rmq *shared.RabbitMqConsumer) {
	defer rmq.Close()

	msgs, err := rmq.Consume()

	if err != nil {
		log.Println(err)
	}

	for msg := range msgs {
		log.Println(string(msg.Body))
	}
}
