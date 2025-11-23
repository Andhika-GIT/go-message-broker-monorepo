package order

import (
	"fmt"
	"net/http"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/worker"
)

type OrderHandler struct {
	usecase      *OrderUseCase
	uploadWorker *worker.UploadWorker
}

func NewOrderHandler(usecase *OrderUseCase, uploadWorker *worker.UploadWorker) *OrderHandler {
	return &OrderHandler{
		usecase:      usecase,
		uploadWorker: uploadWorker,
	}
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	paginationReq := shared.GetPaginationParams(r)

	orderFilter := BindOrderFilterFromRequest(r)

	orders, err := h.usecase.FindAllOrders(r.Context(), paginationReq, orderFilter)

	if err != nil {
		shared.SendJsonErrorResponse(w, err, nil)
		return
	}

	shared.SendJsonResponse(w, 200, "success", orders)

}

func (h *OrderHandler) UploadOrder(w http.ResponseWriter, r *http.Request) {
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
