package user

import (
	"net/http"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

type UserHandler struct {
	usecase *UserUseCase
}

func NewUserHandler(usecase *UserUseCase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.usecase.FindAllUsers(r.Context())

	if err != nil {
		shared.SendJsonErrorResponse(w, err, nil)
		return
	}

	shared.SendJsonResponse(w, 200, "success", users)
}

func (h *UserHandler) UploadUser(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	err := h.usecase.ReadFile(r)

	if err != nil {
		shared.SendJsonErrorResponse(w, err, nil)
		return
	}

	shared.SendJsonResponse(w, 200, "success", nil)
}
