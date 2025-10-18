package user

import (
	"net/http"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

type UserHandler struct {
	rmq     *shared.RabbitMqProducer
	usecase *UserUseCase
}

func NewUserHandler(rmq *shared.RabbitMqProducer, usecase *UserUseCase) *UserHandler {
	return &UserHandler{
		rmq:     rmq,
		usecase: usecase,
	}
}

// func (h *UserHandler)

func (h *UserHandler) UploadUser(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	err := h.usecase.ReadFile(r)

	if err != nil {
		shared.SendJsonErrorResponse(w, err, nil)
		return
	}

	shared.SendJsonResponse(w, 200, "success", nil)
}
