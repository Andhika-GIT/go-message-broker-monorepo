package order

import (
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/rabbitmq/amqp091-go"
)

func StartWorker(rmq *shared.RabbitMqProducer) {
	defer rmq.Close()

	q, err := rmq.Channel.QueueDeclare(
		"test_queue", true, false, false, false,
		amqp091.Table{
			"x-queue-type": "quorum",
		},
	)

	if err != nil {
		log.Fatalf("failed to decleare queue : %v", err.Error())
	}

	err = rmq.Channel.QueueBind(
		q.Name, "test", "go-exchange", false, nil,
	)

	if err != nil {
		log.Fatalf("Failed to bind queue: %v", err)
	}

	msgs, err := rmq.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Fatalf("Failed to consume message: %v", err)
	}

	for msg := range msgs {
		log.Println(string(msg.Body))
	}
}
