package order

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/user"
	"gorm.io/gorm"
)

func NewOrderModule(rmq *shared.RabbitMqConsumer, DB *gorm.DB, userUseCase *user.UserUseCase) {
	orderUseCase := NewOrderUseCase(&OrderRepository{}, DB, userUseCase)

	directUC := NewOrderDirectWorker(rmq, orderUseCase)

	go directUC.Start()

}
