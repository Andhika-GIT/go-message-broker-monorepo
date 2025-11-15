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

func (r *UserRepository) FindAll(c context.Context, paginationReq *shared.PaginationRequest, filter *UserFilter) (*shared.Paginated[User], error) {
	var users []User
	var totalRecords int64

	offset := (paginationReq.Page - 1) * paginationReq.PerPage

	baseQuery := r.DB.WithContext(c).Model(&User{})

	query := FilterUserQuery(filter, baseQuery)

	err := query.Session(&gorm.Session{}).Count(&totalRecords).Error

	if err != nil {
		return nil, err
	}

	err = query.Offset(offset).Limit(paginationReq.PerPage).Find(&users).Error

	totalPages := (int(totalRecords) + paginationReq.PerPage - 1) / paginationReq.PerPage

	if err != nil {
		return nil, err
	}

	return &shared.Paginated[User]{
		Data:       users,
		Total:      totalRecords,
		TotalPages: totalPages,
	}, nil
}
