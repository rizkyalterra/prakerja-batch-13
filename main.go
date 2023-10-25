package main

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	InitDB()
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetDetailUsersController)
	e.POST("/users", AddUsersController)
	e.Start(":8000")
}

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
	DB.AutoMigrate(&User{})
}

func AddUsersController(c echo.Context) error {
	var user User
	c.Bind(&user)

	//

	result := DB.Create(&user)
	if result.Error != nil {
		return c.JSON(500, BaseResponse{
			Status: false,
			Message: "Failed create in database",
			Data: nil,
		})
	}

	return c.JSON(200, BaseResponse{
		Status: true,
		Message: "Successfully created",
		Data: user,
	})
}

func GetDetailUsersController(c echo.Context) error {
	// id := c.Param("id")
	var user User = User{}
	return c.JSON(200, user)
}

func GetUsersController(c echo.Context) error {
	var users []User
	result := DB.Find(&users)
	if result.Error != nil {
		return c.JSON(500, BaseResponse{
			Status: false,
			Message: "Failed get data from database",
			Data: nil,
		})
	}
	return c.JSON(200,  BaseResponse{
		Status: true,
		Message: "Successfully get data",
		Data: users,
	})
}
