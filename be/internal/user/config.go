package user

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/worker"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewUserModule(r chi.Router, rmq *shared.RabbitMqProducer, uploadWorker *worker.UploadWorker,
	DB *gorm.DB, mqRoutingKey *shared.RabbitMQRoutingKey) {
	repository := NewUserRepository(DB)
	usecase := NewUserUseCase(repository, rmq, DB)
	handler := NewUserHandler(usecase, uploadWorker, mqRoutingKey)
	NewUserRoutes(r, handler)
}
