package order

import "gorm.io/gorm"

type OrderUseCase struct {
	Repository *OrderRepository
	DB         *gorm.DB
}

func NewOrderUseCase(Repository *OrderRepository, DB *gorm.DB) *OrderUseCase {
	return &OrderUseCase{
		Repository: Repository,
		DB:         DB,
	}
}
