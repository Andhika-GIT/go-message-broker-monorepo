package user

import (
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

type UserUseCase struct {
	Repository *UserRepository
	DB         *gorm.DB
}

func NewUserUseCase(Repository *UserRepository, DB *gorm.DB) *UserUseCase {
	return &UserUseCase{
		Repository: Repository,
		DB:         DB,
	}
}

func (uc *UserUseCase) CreateNewUsers(c context.Context, ch <-chan UserImport) error {

	for user := range ch {
		tx := uc.DB.WithContext(c).Begin()

		err := uc.Repository.FindByEmail(c, tx, &User{}, user.Email)

		// if user email already exist, throw error
		if err == nil {
			tx.Rollback()
			log.Printf("user already exist")
			continue
		}

		// other error besides not found from Repository.FindByEmail
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			log.Printf("unexpected error: %v", err)
			continue
		}

		err = uc.Repository.Create(c, tx, &User{
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		})

		if err != nil {
			tx.Rollback()
			log.Printf("error : %s", err.Error())
			continue
		}

		err = tx.Commit().Error

		if err != nil {
			log.Printf("error : %s", err.Error())
		}
	}

	return nil
}

func (uc *UserUseCase) FindUserByEmail(c context.Context, userEmail string) (UserResponse, error) {
	var user User

	tx := uc.DB.WithContext(c)

	err := uc.Repository.FindByEmail(c, tx, &user, userEmail)

	if err != nil {
		return UserResponse{}, err
	}

	return UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}, nil
}
