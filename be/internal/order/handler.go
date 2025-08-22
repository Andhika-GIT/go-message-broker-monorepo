package order

import (
	"log"
	"net/http"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/rabbitmq/amqp091-go"
)

type OrderHandler struct {
	rmq *shared.RabbitMqProducer
}

func NewOrderHandler(rmq *shared.RabbitMqProducer) *OrderHandler {
	return &OrderHandler{
		rmq: rmq,
	}
}

func (h *OrderHandler) TestRabbitMq(w http.ResponseWriter, r *http.Request) {
	q, err := h.rmq.Channel.QueueDeclare(
		"test_queue", true, false, false, false,
		amqp091.Table{
			"x-queue-type": "quorum",
		},
	)

	if err != nil {
		log.Fatalf("failed to decleare queue : %v", err.Error())
	}

	err = h.rmq.Channel.QueueBind(
		q.Name, "test", "go-exchange", false, nil,
	)

	if err != nil {
		log.Fatalf("Failed to bind queue: %v", err)
	}

	err = h.rmq.Channel.Publish(
		"go-exchange", "test", false, false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte("ini dari handler controller"),
		},
	)

	if err != nil {
		log.Fatalf("failed to publish message: %v", err)
	}
}
