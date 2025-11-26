package user

import (
	"encoding/json"
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/pkg/sftp"
)

type UserDirectUploadWorker struct {
	Rmq        *shared.RabbitMqConsumer
	UseCase    *UserUseCase
	QueueCfg   *shared.RabbitMQQueue
	sftpClient *sftp.Client
}

func NewUserDirectUploadWorker(Rmq *shared.RabbitMqConsumer, UseCase *UserUseCase, cfg *shared.RabbitMQQueue, sftpClient *sftp.Client) *UserDirectUploadWorker {
	return &UserDirectUploadWorker{
		Rmq:        Rmq,
		UseCase:    UseCase,
		QueueCfg:   cfg,
		sftpClient: sftpClient,
	}
}

func (w *UserDirectUploadWorker) Start() {
	defer w.Rmq.Close()

	msgs, err := w.Rmq.Consume(w.QueueCfg.UserDirectImport)

	if err != nil {
		log.Println(err)
	}

	var uploadMsg shared.UploadMessage
	for msg := range msgs {
		err := json.Unmarshal(msg.Body, &uploadMsg)

		if err != nil {
			log.Println("error when converting", err.Error())
			continue
		}

		remoteFile, err := w.sftpClient.Open(uploadMsg.Filepath)

		if err != nil {
			log.Fatalf("error when reading sftp file: %v", err)
			continue
		}

		rows, err := shared.ReadExcel(remoteFile)

		if err != nil {
			log.Fatal(err.Error())
		}

		log.Printf("rows are %v", rows)

	}

}
