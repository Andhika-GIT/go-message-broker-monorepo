package user

import (
	"context"
	"fmt"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/pkg/sftp"
	"gorm.io/gorm"
)

type UserUseCase struct {
	Repository *UserRepository
	rmq        *shared.RabbitMqProducer
	DB         *gorm.DB
	sftp       *sftp.Client
}

func NewUserUseCase(Repository *UserRepository, rmq *shared.RabbitMqProducer, sftp *sftp.Client, DB *gorm.DB) *UserUseCase {
	return &UserUseCase{
		Repository: Repository,
		rmq:        rmq,
		DB:         DB,
		sftp:       sftp,
	}
}

func (u *UserUseCase) FindAllUsers(c context.Context, paginationReq *shared.PaginationRequest, filter *UserFilter) (*shared.Paginated[UserResponse], error) {

	paginated, err := u.Repository.FindAll(c, paginationReq, filter)

	if err != nil {
		return nil, shared.WriteError(500, fmt.Sprintf("failed to find all users %s", err.Error()))
	}

	formatedUsers := ConvertToUsersResponse(paginated.Data)

	// return new paginated response with different type (UserResponse)
	return &shared.Paginated[UserResponse]{
		Data:       formatedUsers,
		Total:      paginated.Total,
		TotalPages: paginated.TotalPages,
	}, nil

}
