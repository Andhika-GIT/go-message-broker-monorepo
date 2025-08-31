package user

import (
	"fmt"
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

func (h *UserHandler) UploadUser(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("file")

	if err != nil {
		shared.SendJsonResponse(w, 500, fmt.Sprintf("failed to read file %s", err.Error()), nil)
		return
	}

	defer file.Close()

	isAllowedExtension := shared.IsAllowedExtension(header.Filename)

	if !isAllowedExtension {
		shared.SendJsonResponse(w, 400, "invalid file extension", nil)
		return
	}

	err = h.usecase.ReadFile(file)

	if err != nil {
		shared.SendJsonErrorResponse(w, err, nil)
		return
	}

	shared.SendJsonResponse(w, 200, "success", nil)
}
