package order

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	redispubsub "github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared/redis"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/user"
	"github.com/pkg/sftp"
	"gorm.io/gorm"
)

func NewOrderModule(rmq *shared.RabbitMqConsumer, rdsPublisher *redispubsub.Publisher, DB *gorm.DB, userUseCase *user.UserUseCase, queueCfg *shared.RabbitMQQueue, sftpClient *sftp.Client) {
	orderUseCase := NewOrderUseCase(&OrderRepository{}, DB, userUseCase)

	directUC := NewOrderDirectWorker(rmq, orderUseCase, queueCfg, sftpClient)

	go directUC.Start()

}
