package order

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/go-chi/chi/v5"
)

func NewOrderModule(r chi.Router, rmq *shared.RabbitMqProducer) {
	handler := NewOrderHandler(rmq)
	NewOrderRoutes(r, handler)
}
