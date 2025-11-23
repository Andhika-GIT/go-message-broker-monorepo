package user

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/worker"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/sftp"
	"gorm.io/gorm"
)

func NewUserModule(r chi.Router, rmq *shared.RabbitMqProducer, sftpClient *sftp.Client, uploadWorker *worker.UploadWorker,
	DB *gorm.DB) {
	repository := NewUserRepository(DB)
	usecase := NewUserUseCase(repository, rmq, sftpClient, DB)
	handler := NewUserHandler(usecase, uploadWorker)
	NewUserRoutes(r, handler)
}
