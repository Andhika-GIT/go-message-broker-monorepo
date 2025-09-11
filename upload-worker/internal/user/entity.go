package user

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"primary_key;column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Password  string    `gorm:"column:password; not null"`
	Email     string    `json:"email" gorm:"column:email"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime"`
}

func (a *User) TableName() string {
	return "users"
}
