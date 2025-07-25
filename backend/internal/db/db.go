package db

import (
	"glasscutting/internal/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Init() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=glasscutting port=5432 sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{}, &model.Order{}, &model.Service{})
	return db
}
