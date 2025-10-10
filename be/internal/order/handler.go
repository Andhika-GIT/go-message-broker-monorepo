package order

import (
	"net/http"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

type OrderHandler struct {
	rmq     *shared.RabbitMqProducer
	usecase *OrderUseCase
}

func NewOrderHandler(rmq *shared.RabbitMqProducer, usecase *OrderUseCase) *OrderHandler {
	return &OrderHandler{
		rmq:     rmq,
		usecase: usecase,
	}
}

func (h *OrderHandler) TestRabbitMq(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	err := h.usecase.ReadFile(r)

	if err != nil {
		shared.SendJsonErrorResponse(w, err, nil)
	}

	shared.SendJsonResponse(w, 200, "success", nil)
}
