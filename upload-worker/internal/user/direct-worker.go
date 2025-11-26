package user

import (
	"encoding/json"
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

type UserDirectUploadWorker struct {
	Rmq      *shared.RabbitMqConsumer
	UseCase  *UserUseCase
	QueueCfg *shared.RabbitMQQueue
}

func NewUserDirectUploadWorker(Rmq *shared.RabbitMqConsumer, UseCase *UserUseCase, cfg *shared.RabbitMQQueue) *UserDirectUploadWorker {
	return &UserDirectUploadWorker{
		Rmq:      Rmq,
		UseCase:  UseCase,
		QueueCfg: cfg,
	}
}

func (w *UserDirectUploadWorker) Start() {
	defer w.Rmq.Close()

	msgs, err := w.Rmq.Consume(w.QueueCfg.UserDirectImport)

	if err != nil {
		log.Println(err)
	}

	var uploadmsg shared.UploadMessage
	for msg := range msgs {
		err := json.Unmarshal(msg.Body, &uploadmsg)

		if err != nil {
			log.Println("error when converting", err.Error())
			continue
		}

		log.Printf("message is %v", uploadmsg)

	}

}
