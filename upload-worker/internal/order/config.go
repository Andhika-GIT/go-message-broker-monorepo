package order

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"gorm.io/gorm"
)

func NewOrderModule(rmq *shared.RabbitMqConsumer, DB *gorm.DB) {
	orderUseCase := NewOrderUseCase(&OrderRepository{}, DB)

	directUC := NewOrderDirectWorker(rmq, orderUseCase)

	go directUC.Start()

}
