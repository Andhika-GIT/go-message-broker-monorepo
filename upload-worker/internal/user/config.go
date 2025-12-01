package user

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	redispubsub "github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared/redis"
	"github.com/pkg/sftp"
	"gorm.io/gorm"
)

type UserModule struct {
	UserUseCase *UserUseCase
}

func NewUserModule(rmq *shared.RabbitMqConsumer, rdsPublisher *redispubsub.Publisher, DB *gorm.DB, queueCfg *shared.RabbitMQQueue, sftpClient *sftp.Client) *UserModule {
	userUseCase := NewUserUseCase(&UserRepository{}, DB)

	directUC := NewUserDirectUploadWorker(rmq, rdsPublisher, userUseCase, queueCfg, sftpClient)

	go directUC.Start()

	return &UserModule{
		UserUseCase: userUseCase,
	}
}
