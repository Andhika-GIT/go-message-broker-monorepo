package order

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/user"
	"gorm.io/gorm"
)

type OrderUseCase struct {
	Repository  *OrderRepository
	DB          *gorm.DB
	UserUseCase *user.UserUseCase
}

func stringToOrderStatus(statusStr string) OrderStatus {
	switch strings.ToLower(strings.TrimSpace(statusStr)) {
	case "pending", "Pending":
		return StatusPending
	case "processing", "process", "Processing":
		return StatusProcessing
	case "completed", "complete", "Completed":
		return StatusCompleted
	case "cancelled", "cancel", "Cancelled":
		return StatusCancelled
	default:
		return StatusPending
	}
}

func NewOrderUseCase(Repository *OrderRepository, DB *gorm.DB, UserUseCase *user.UserUseCase) *OrderUseCase {
	return &OrderUseCase{
		Repository:  Repository,
		DB:          DB,
		UserUseCase: UserUseCase,
	}
}

func (uc *OrderUseCase) ReadOrderExcel(rows [][]string) []OrderImport {
	var orders []OrderImport

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) >= 4 {
			quantity, err := strconv.Atoi(row[2])
			if err != nil {
				continue
			}

			status := stringToOrderStatus(row[3])

			orders = append(orders, OrderImport{
				UserEmail:   row[0],
				ProductName: row[1],
				Quantity:    quantity,
				Status:      status,
			})
		}
	}

	return orders
}

func (uc *OrderUseCase) CreateOrders(c context.Context, orders []OrderImport) error {
	tx := uc.DB.WithContext(c).Begin()

	defer tx.Rollback()

	var newOrders []Order

	for _, order := range orders {
		user, err := uc.UserUseCase.FindUserByEmail(c, order.UserEmail)
		if err != nil {
			log.Printf("failed to find user: %s", err.Error())
			continue
		}

		newOrders = append(newOrders, Order{
			UserId:      user.ID,
			ProductName: order.ProductName,
			Quantity:    int64(order.Quantity),
			Status:      string(order.Status),
		})

	}

	err := uc.Repository.Create(c, tx, &newOrders)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
