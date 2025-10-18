package order

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/go-chi/chi/v5"
)

func NewOrderModule(r chi.Router, rmq *shared.RabbitMqProducer) {
	usecase := NewOrderUseCase(rmq)
	handler := NewOrderHandler(usecase)
	NewOrderRoutes(r, handler)
}
