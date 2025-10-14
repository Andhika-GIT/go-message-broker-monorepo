package user

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"gorm.io/gorm"
)

type UserModule struct {
	UserUseCase *UserUseCase
}

func NewUserModule(rmq *shared.RabbitMqConsumer, DB *gorm.DB) *UserModule {
	userUseCase := NewUserUseCase(&UserRepository{}, DB)

	directUC := NewUserDirectUploadWorker(rmq, userUseCase)

	go directUC.Start()

	return &UserModule{
		UserUseCase: userUseCase,
	}
}
