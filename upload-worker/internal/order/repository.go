package order

import (
	"context"

	"gorm.io/gorm"
)

type OrderRepository struct{}

func (r *OrderRepository) Create(c context.Context, tx *gorm.DB, order *Order) error {
	err := tx.Create(&order).Error

	if err != nil {
		return err
	}

	return nil
}
