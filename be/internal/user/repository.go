package user

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository struct{}

func (r *UserRepository) FindAll(c context.Context, tx *gorm.DB, users *[]User) error {
	err := tx.Find(&users).Error

	if err != nil {
		return err
	}

	return nil
}
