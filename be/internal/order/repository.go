package order

import (
	"context"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: DB,
	}
}

func (r *OrderRepository) CountAll(c context.Context) (*int64, error) {
	var totalRecords *int64

	baseQuery := r.DB.WithContext(c).Model(&Order{})

	err := baseQuery.Count(totalRecords).Error

	if err != nil {
		return nil, err
	}

	return totalRecords, nil

}

func (r *OrderRepository) FindAll(c context.Context, paginationReq *shared.PaginationRequest, filter *OrderFilter) (*shared.Paginated[Order], error) {
	var orders []Order
	var totalRecords int64

	offset := (paginationReq.Page - 1) * paginationReq.PerPage

	baseQuery := r.DB.WithContext(c).Model(&Order{}).Preload("User")

	query := FilterOrderQuery(filter, baseQuery)

	err := query.Session(&gorm.Session{}).Count(&totalRecords).Error

	if err != nil {
		return nil, err
	}

	err = baseQuery.Offset(offset).Limit(paginationReq.PerPage).Find(&orders).Error

	if err != nil {
		return nil, err
	}

	totalPages := (int(totalRecords) + paginationReq.PerPage - 1) / paginationReq.PerPage

	return &shared.Paginated[Order]{
		Data:       orders,
		Total:      totalRecords,
		TotalPages: totalPages,
	}, nil
}
