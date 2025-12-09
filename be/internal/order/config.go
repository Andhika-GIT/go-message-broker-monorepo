package order

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/worker"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type OrderModule struct {
	UseCase *OrderUseCase
}

func NewOrderModule(r chi.Router, rmq *shared.RabbitMqProducer, uploadWorker *worker.UploadWorker, DB *gorm.DB, cfg *shared.Config) *OrderModule {
	repository := NewOrderRepository(DB)
	usecase := NewOrderUseCase(repository, rmq)
	handler := NewOrderHandler(usecase, uploadWorker, &cfg.RabbitMQRoutingKey, cfg.SftpClient.Path)
	NewOrderRoutes(r, handler)

	return &OrderModule{
		UseCase: usecase,
	}
}
