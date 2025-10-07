package user

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"gorm.io/gorm"
)

func NewUserModule(rmq *shared.RabbitMqConsumer, DB *gorm.DB) {
	userUseCase := NewUserUseCase(&UserRepository{}, DB)

	directUC := NewUserDirectUploadWorker(rmq, userUseCase)

	go directUC.Start()
}
