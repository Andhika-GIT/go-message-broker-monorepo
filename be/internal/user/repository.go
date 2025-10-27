package user

import (
	"context"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: DB,
	}
}

func (r *UserRepository) FindAll(c context.Context, paginationReq *shared.PaginationRequest) ([]User, error) {
	var users []User

	offset := (paginationReq.Page - 1) * paginationReq.PerPage

	query := r.DB.Offset(offset).Limit(paginationReq.PerPage)

	err := query.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}
