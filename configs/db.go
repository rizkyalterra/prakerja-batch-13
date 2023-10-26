package configs

import (
	"prakerja13/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja13?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect database")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&models.User{})
}
