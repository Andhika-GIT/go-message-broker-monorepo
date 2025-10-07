package order

import "github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"

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

func (w *OrderDirectWorker) Start() {}
