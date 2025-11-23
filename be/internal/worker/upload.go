package worker

import (
	"fmt"
	"io"
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/pkg/sftp"
)

type UploadTask struct {
	File     io.Reader
	FileName string
}

var UploadQueue = make(chan UploadTask, 100)

type UploadWorker struct {
	rmq        *shared.RabbitMqProducer
	sftpClient *sftp.Client
	numWorkers int
}

func NewUploadWorker(sftp *sftp.Client, rmq *shared.RabbitMqProducer, numWorkers int) *UploadWorker {
	return &UploadWorker{
		sftpClient: sftp,
		rmq:        rmq,
		numWorkers: numWorkers,
	}
}

func (w *UploadWorker) Start() {
	for i := 0; i < w.numWorkers; i++ {
		go func() {
			for task := range UploadQueue {
				w.processUpload(task)
			}
		}()
	}
}

func (w *UploadWorker) processUpload(task UploadTask) {
	dstFile, err := w.sftpClient.Create(fmt.Sprintf("/upload/%s", task.FileName))

	if err != nil {
		log.Fatalf("error when connecting to sftp : %s", err.Error())
	}

	defer dstFile.Close()

	_, err = io.Copy(dstFile, task.File)

	if err != nil {
		log.Fatalf("error when insert file to sftp %s", err.Error())
	}

}

func (w *UploadWorker) Queue(task UploadTask) {
	UploadQueue <- task
}
