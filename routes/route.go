package routes

import (
	"prakerja13/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo){
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetDetailUsersController)
	e.POST("/users", controllers.AddUsersController)
}