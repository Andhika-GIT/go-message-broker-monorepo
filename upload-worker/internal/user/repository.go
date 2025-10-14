package user

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository struct{}

func (r *UserRepository) Create(c context.Context, tx *gorm.DB, user *User) error {
	err := tx.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByEmail(c context.Context, tx *gorm.DB, user *User, userEmail string) error {
	err := tx.Where("email = ?", userEmail).First(&user).Error

	if err != nil {
		return err
	}

	return nil
}
