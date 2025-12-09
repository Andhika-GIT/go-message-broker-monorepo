package dashboard

import (
	"net/http"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/order"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/user"
)

type DashboardHandler struct {
	userUseCase  *user.UserUseCase
	orderUseCase *order.OrderUseCase
}

func NewDashboardHandler(userUseCase *user.UserUseCase, orderUseCase *order.OrderUseCase) *DashboardHandler {
	return &DashboardHandler{
		userUseCase:  userUseCase,
		orderUseCase: orderUseCase,
	}
}

func (h *DashboardHandler) GetDataSummary(w http.ResponseWriter, r *http.Request) {

}
