package user

import (
	"encoding/json"
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

type DirectUploadWorker struct {
	Rmq     *shared.RabbitMqConsumer
	UseCase *UserUseCase
}

func NewDirectUploadWorker(Rmq *shared.RabbitMqConsumer, UseCase *UserUseCase) *DirectUploadWorker {
	return &DirectUploadWorker{
		Rmq:     Rmq,
		UseCase: UseCase,
	}
}

func (w *DirectUploadWorker) Start() {
	defer w.Rmq.Close()

	msgs, err := w.Rmq.Consume(shared.QueueUserDirectImport)

	if err != nil {
		log.Println(err)
	}

	var users []UserImport
	for msg := range msgs {
		err := json.Unmarshal(msg.Body, &users)

		if err != nil {
			log.Println("error when converting", err.Error())
			continue
		}

		log.Printf("users is %v", users)

		for _, user := range users {
			log.Printf("Name: %s, Email: %s, Phone: %s\n", user.Name, user.Email, user.PhoneNumber)

		}
	}

}
