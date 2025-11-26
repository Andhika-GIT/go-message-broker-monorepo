package user

import (
	"fmt"
	"net/http"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/worker"
)

type UserHandler struct {
	usecase      *UserUseCase
	uploadWorker *worker.UploadWorker
	mqRoutingKey *shared.RabbitMQRoutingKey
}

func NewUserHandler(usecase *UserUseCase, uploadWorker *worker.UploadWorker, mqRoutingKey *shared.RabbitMQRoutingKey) *UserHandler {
	return &UserHandler{
		usecase:      usecase,
		uploadWorker: uploadWorker,
		mqRoutingKey: mqRoutingKey,
	}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	paginationReq := shared.GetPaginationParams(r)

	userFilter := BindUserFilterFromRequest(r)

	users, err := h.usecase.FindAllUsers(r.Context(), paginationReq, userFilter)

	if err != nil {
		shared.SendJsonErrorResponse(w, err, nil)
		return
	}

	shared.SendJsonResponse(w, 200, "success", users)
}

func (h *UserHandler) UploadUser(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("file")

	if err != nil {
		shared.WriteError(500, fmt.Sprintf("failed to read file %s", err.Error()))
		return
	}

	defer file.Close()

	isFileExtensionCorrect := shared.IsAllowedExtension(header.Filename)

	if !isFileExtensionCorrect {
		shared.WriteError(400, "invalid file extension")
		return
	}

	h.uploadWorker.Queue(worker.UploadTask{
		File:     file,
		FileName: header.Filename,
	})

	shared.SendJsonResponse(w, 200, "success", nil)
}
