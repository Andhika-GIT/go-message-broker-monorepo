package user

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/worker"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type UserModule struct {
	useCase UserUseCase
}

func NewUserModule(r chi.Router, rmq *shared.RabbitMqProducer, uploadWorker *worker.UploadWorker,
	DB *gorm.DB, cfg *shared.Config) *UserModule {
	repository := NewUserRepository(DB)
	usecase := NewUserUseCase(repository, rmq, DB)
	handler := NewUserHandler(usecase, uploadWorker, &cfg.RabbitMQRoutingKey, cfg.SftpClient.Path)
	NewUserRoutes(r, handler)

	return &UserModule{
		useCase: *usecase,
	}
}
