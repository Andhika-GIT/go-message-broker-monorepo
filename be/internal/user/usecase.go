package user

import "github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"

type UserUseCase struct {
	rmq     *shared.RabbitMqProducer
	handler *UserHandler
}

func NewUserUseCase(rmq *shared.RabbitMqProducer, handler *UserHandler) *UserUseCase {
	return &UserUseCase{
		rmq:     rmq,
		handler: handler,
	}
}
