package repositories

import (
	"fmt"
	"prakerja13/configs"
	"prakerja13/models"
)

func AddUser(user *models.User) error {
	result := configs.DB.Create(user)
	return result.Error
}

func GetUser(users *[]models.User) error {
	result := configs.DB.Find(users)
	fmt.Println(users)
	return result.Error
}
