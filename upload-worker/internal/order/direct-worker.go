package order

import (
	"encoding/json"
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

type OrderDirectWorker struct {
	Rmq     *shared.RabbitMqConsumer
	UseCase *OrderUseCase
}

func NewOrderDirectWorker(Rmq *shared.RabbitMqConsumer, UseCase *OrderUseCase) *OrderDirectWorker {
	return &OrderDirectWorker{
		Rmq:     Rmq,
		UseCase: UseCase,
	}
}

func (w *OrderDirectWorker) Start() {
	defer w.Rmq.Close()

	// ch := make(chan OrderImport)
	// c := context.Background()

	msgs, err := w.Rmq.Consume(shared.QueueOrderDirectImport)

	if err != nil {
		log.Println(err.Error())
	}

	var orders []OrderImport
	for msg := range msgs {
		err := json.Unmarshal(msg.Body, &orders)

		if err != nil {
			log.Panicln(err.Error())
			continue
		}

		for _, order := range orders {
			log.Printf("Name: %s, Email: %s, Phone: %s\n", order.Email, order.ProductName, order.Quantity)
		}
	}
}
