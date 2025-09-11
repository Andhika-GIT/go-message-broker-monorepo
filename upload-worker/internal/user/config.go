package user

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

func NewUserModule(rmq *shared.RabbitMqConsumer) {
	userUseCase := NewUserUseCase(&UserRepository{})

	directUC := NewDirectUploadWorker(rmq, userUseCase)

	go directUC.Start()
}
