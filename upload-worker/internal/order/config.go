package order

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/user"
	"github.com/pkg/sftp"
	"gorm.io/gorm"
)

func NewOrderModule(rmq *shared.RabbitMqConsumer, DB *gorm.DB, userUseCase *user.UserUseCase, queueCfg *shared.RabbitMQQueue, sftpClient *sftp.Client) {
	orderUseCase := NewOrderUseCase(&OrderRepository{}, DB, userUseCase)

	directUC := NewOrderDirectWorker(rmq, orderUseCase, queueCfg, sftpClient)

	go directUC.Start()

}
