package main

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()
	e.GET("/users", GetUsersController)
    e.GET("/users/:id", GetDetailUsersController)
    e.POST("/users", AddUsersController)
	e.Start(":8000")
}

func AddUsersController(c echo.Context) error {
    var user User
	c.Bind(&user)
	return c.JSON(200, user)
}

func GetDetailUsersController(c echo.Context) error {
    id := c.Param("id")
	var user User = User{1, id, "alta@gmail.com"}
	return c.JSON(200, user)
}

func GetUsersController(c echo.Context) error {
    country := c.QueryParam("country")
	var user User = User{1, country, "alta@gmail.com"}
	return c.JSON(200, user)
}
