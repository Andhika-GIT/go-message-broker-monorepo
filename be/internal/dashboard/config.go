package dashboard

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/order"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/user"
	"github.com/go-chi/chi/v5"
)

func NewDashboardModule(r chi.Router, userUseCase *user.UserUseCase, orderUseCase *order.OrderUseCase) {
	handler := NewDashboardHandler(userUseCase, orderUseCase)
	NewDashboardRoutes(r, handler)

}
