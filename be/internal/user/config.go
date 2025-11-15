package user

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewUserModule(r chi.Router, rmq *shared.RabbitMqProducer, DB *gorm.DB) {
	repository := NewUserRepository(DB)
	usecase := NewUserUseCase(repository, rmq, DB)
	handler := NewUserHandler(usecase)
	NewUserRoutes(r, handler)
}
