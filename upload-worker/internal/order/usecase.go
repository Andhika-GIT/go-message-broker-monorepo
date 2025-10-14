package order

import (
	"context"
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/user"
	"gorm.io/gorm"
)

type OrderUseCase struct {
	Repository  *OrderRepository
	DB          *gorm.DB
	UserUseCase *user.UserUseCase
}

func NewOrderUseCase(Repository *OrderRepository, DB *gorm.DB, UserUseCase *user.UserUseCase) *OrderUseCase {
	return &OrderUseCase{
		Repository:  Repository,
		DB:          DB,
		UserUseCase: UserUseCase,
	}
}

func (uc *OrderUseCase) CreateOrders(c context.Context, ch <-chan OrderImport) {
	for order := range ch {
		tx := uc.DB.WithContext(c).Begin()

		user, err := uc.UserUseCase.FindUserByEmail(c, order.UserEmail)

		if err != nil {
			log.Printf("failed to find user : %s", err.Error())
			tx.Rollback()
			continue
		}

		err = uc.Repository.Create(c, tx, &Order{
			UserId:      user.ID,
			ProductName: order.ProductName,
			Quantity:    int64(order.Quantity),
			Status:      string(order.Status),
		})

		if err != nil {
			log.Printf("failed to save order : %s", err.Error())
			tx.Rollback()
			continue
		}

		err = tx.Commit().Error

		if err != nil {
			log.Printf("error : %s", err.Error())
		}

	}

}
