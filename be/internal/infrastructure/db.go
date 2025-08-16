package infrastructure

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(viper *viper.Viper) *gorm.DB {
	DB_HOST := viper.GetString("DB_HOST")
	DB_NAME := viper.GetString("DB_NAME")
	DB_PORT := viper.GetString("DB_PORT")
	DB_USERNAME := viper.GetString("DB_USERNAME")
	DB_PASSWORD := viper.GetString("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("fatal error connecting to database: %w", err))
	}

	return db

}
