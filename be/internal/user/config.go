package user

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/go-chi/chi/v5"
)

func NewUserModule(r chi.Router, rmq *shared.RabbitMqProducer) {
	handler := NewUserHandler(rmq)
	NewUserRoutes(r, handler)
}
