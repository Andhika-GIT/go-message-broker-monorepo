package order

import (
	"net/http"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

type OrderHandler struct {
	usecase *OrderUseCase
}

func NewOrderHandler(usecase *OrderUseCase) *OrderHandler {
	return &OrderHandler{
		usecase: usecase,
	}
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	paginationReq := shared.GetPaginationParams(r)

	orders, err := h.usecase.FindAllOrders(r.Context(), paginationReq)

	if err != nil {
		shared.SendJsonErrorResponse(w, err, nil)
		return
	}

	shared.SendJsonResponse(w, 200, "success", orders)

}

func (h *OrderHandler) UploadOrder(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	err := h.usecase.ReadFile(r)

	if err != nil {
		shared.SendJsonErrorResponse(w, err, nil)
		return
	}

	shared.SendJsonResponse(w, 200, "success", nil)
}
