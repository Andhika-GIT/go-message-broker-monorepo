package order

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewOrderModule(r chi.Router, rmq *shared.RabbitMqProducer, DB *gorm.DB) {
	repository := NewOrderRepository(DB)
	usecase := NewOrderUseCase(repository, rmq)
	handler := NewOrderHandler(usecase)
	NewOrderRoutes(r, handler)
}
