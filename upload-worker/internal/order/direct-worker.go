package order

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

type OrderDirectWorker struct {
	Rmq      *shared.RabbitMqConsumer
	UseCase  *OrderUseCase
	QueueCfg *shared.RabbitMQQueue
}

func NewOrderDirectWorker(Rmq *shared.RabbitMqConsumer, UseCase *OrderUseCase, cfg *shared.RabbitMQQueue) *OrderDirectWorker {
	return &OrderDirectWorker{
		Rmq:      Rmq,
		UseCase:  UseCase,
		QueueCfg: cfg,
	}
}

func (w *OrderDirectWorker) Start() {
	defer w.Rmq.Close()

	ch := make(chan OrderImport)
	c := context.Background()

	msgs, err := w.Rmq.Consume(w.QueueCfg.OrderDirectImport)

	if err != nil {
		log.Println(err.Error())
	}

	for i := 0; i <= 3; i++ {
		go w.UseCase.CreateOrders(c, ch)
	}

	var orders []OrderImport
	for msg := range msgs {
		err := json.Unmarshal(msg.Body, &orders)

		if err != nil {
			log.Panicln(err.Error())
			continue
		}

		for _, order := range orders {
			ch <- order
		}
	}
}
