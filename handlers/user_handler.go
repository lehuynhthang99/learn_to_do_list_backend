package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lehuynhthang99/learn_to_do_list_backend/config"
	"github.com/lehuynhthang99/learn_to_do_list_backend/middleware"
	"github.com/lehuynhthang99/learn_to_do_list_backend/models"
)

// POST /users
func CreateUser(c echo.Context) error {
	customContext, _ := middleware.GetCustomContext(c)

	user := new(models.User)
	err := customContext.Bind(&user)
	if err != nil {
		return customContext.JSON(http.StatusBadRequest, echo.Map{
			"error":   "Invalid input",
			"message": err.Error(),
		})
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		return customContext.JSON(http.StatusInternalServerError, echo.Map{
			"error": result.Error.Error(),
		})
	}

	return customContext.JSON(http.StatusCreated, user)
}
