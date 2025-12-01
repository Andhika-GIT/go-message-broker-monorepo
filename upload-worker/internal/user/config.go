package user

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/pkg/sftp"
	"gorm.io/gorm"
)

type UserModule struct {
	UserUseCase *UserUseCase
}

func NewUserModule(rmq *shared.RabbitMqConsumer, DB *gorm.DB, queueCfg *shared.RabbitMQQueue, sftpClient *sftp.Client) *UserModule {
	userUseCase := NewUserUseCase(&UserRepository{}, DB)

	directUC := NewUserDirectUploadWorker(rmq, userUseCase, queueCfg, sftpClient)

	go directUC.Start()

	return &UserModule{
		UserUseCase: userUseCase,
	}
}
