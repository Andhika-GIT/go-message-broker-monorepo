package user

import (
	"log"
	"net/http"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

type UserHandler struct {
	rmq *shared.RabbitMqProducer
}

func NewUserHandler(rmq *shared.RabbitMqProducer) *UserHandler {
	return &UserHandler{
		rmq: rmq,
	}
}

func (h *UserHandler) UploadUser(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("file")

	if err != nil {
		log.Printf("Failed to read file : %v", err)
	}

	defer file.Close()

	log.Printf("uploaded file : %s (%d bytes)", header.Filename, header.Size)
}
