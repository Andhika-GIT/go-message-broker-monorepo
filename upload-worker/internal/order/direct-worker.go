package order

import (
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

	msgs, err := w.Rmq.Consume(w.QueueCfg.OrderDirectImport)

	if err != nil {
		log.Println(err.Error())
	}

	var uploadMsg shared.UploadMessage
	for msg := range msgs {
		err := json.Unmarshal(msg.Body, &uploadMsg)

		if err != nil {
			log.Panicln(err.Error())
			continue
		}

	}
}
