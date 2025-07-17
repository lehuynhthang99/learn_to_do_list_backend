package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lehuynhthang99/learn_to_do_list_backend/handlers"
	"github.com/lehuynhthang99/learn_to_do_list_backend/middleware"
)

func RegisterRoutes() (e *echo.Echo) {
	e = echo.New()

	//Add middleware
	e.Use(middleware.LoggerMiddleware)
	e.Use(middleware.AuthMiddleWare)

	//user routes
	e.POST("/users", handlers.CreateUser)

	//to do list routes
	e.GET("/to_do_list/:id", handlers.GetAllToDoList)
	e.POST("/to_do_list", handlers.CreateToDoList)
	e.PUT("/to_do_list/:id/items", handlers.UpdateToDoList)
	e.DELETE("/to_do_list/:id", handlers.DeleteToDoList)

	return
}
