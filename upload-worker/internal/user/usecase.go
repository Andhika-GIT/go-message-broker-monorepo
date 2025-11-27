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

func (uc *UserUseCase) ReadUsersExcel(rows [][]string) []UserImport {
	var users []UserImport

	for i, row := range rows {

		if i == 0 {
			continue
		}

		if len(row) >= 3 {
			users = append(users, UserImport{
				Name:        row[0],
				Email:       row[1],
				PhoneNumber: row[2],
			})
		}
	}

	return users
}

func NewUserUseCase(Repository *UserRepository, DB *gorm.DB) *UserUseCase {
	return &UserUseCase{
		Repository: Repository,
		DB:         DB,
	}
}

func (uc *UserUseCase) CreateNewUsers(c context.Context, users []UserImport) error {
	tx := uc.DB.WithContext(c).Begin()

	defer tx.Rollback()

	var newUsers []User

	for _, user := range users {

		err := uc.Repository.FindByEmail(c, tx, &User{}, user.Email)

		// if user email already exist, skip this user
		if err == nil {
			log.Printf("user already exist")
			continue
		}

		// other error besides not found from Repository.FindByEmail
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("unexpected error: %v", err)
			continue
		}

		newUsers = append(newUsers, User{
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		})

	}

	err := uc.Repository.Create(c, tx, &newUsers)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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
