package user

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

type UserDirectUploadWorker struct {
	Rmq     *shared.RabbitMqConsumer
	UseCase *UserUseCase
}

func NewUserDirectUploadWorker(Rmq *shared.RabbitMqConsumer, UseCase *UserUseCase) *UserDirectUploadWorker {
	return &UserDirectUploadWorker{
		Rmq:     Rmq,
		UseCase: UseCase,
	}
}

func (w *UserDirectUploadWorker) Start() {
	defer w.Rmq.Close()

	ch := make(chan UserImport)
	c := context.Background()

	msgs, err := w.Rmq.Consume(shared.QueueUserDirectImport)

	if err != nil {
		log.Println(err)
	}

	for i := 0; i <= 3; i++ {
		// create 3 workers
		go w.UseCase.CreateNewUsers(c, ch)
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
			ch <- user
		}
	}

	close(ch)

}
