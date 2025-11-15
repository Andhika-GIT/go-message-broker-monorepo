package order

import (
	"time"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/user"
)

type Order struct {
	ID          int64     `json:"id" gorm:"primary_key;column:id"`
	UserId      int64     `json:"user_id" gorm:"column:user_id"`
	ProductName string    `json:"product_name" gorm:"column:product_name"`
	Quantity    int64     `json:"quantity" gorm:"column:quantity"`
	Status      string    `json:"status" gorm:"column:status"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime"`

	// user relationship
	User *user.User `json:"user,omitempty" gorm:"foreignKey:UserId;references:ID"`
}

func (a *Order) TableName() string {
	return "orders"
}
