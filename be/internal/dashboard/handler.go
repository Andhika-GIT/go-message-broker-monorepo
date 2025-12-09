package dashboard

import (
	"net/http"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/order"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
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
	totalUsers, err := h.userUseCase.CountAllUsers(r.Context())

	if err != nil {
		shared.SendJsonErrorResponse(w, err, nil)
		return
	}

	totalOrders, err := h.orderUseCase.CountAllOrders(r.Context())

	if err != nil {
		shared.SendJsonErrorResponse(w, err, nil)
		return
	}

	response := &DashboardResponse{
		TotalOrders: *totalOrders,
		TotalUsers:  *totalUsers,
	}

	shared.SendJsonResponse(w, 200, "successfully get dashboard data", response)
}
