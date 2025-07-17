package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lehuynhthang99/learn_to_do_list_backend/config"
	"github.com/lehuynhthang99/learn_to_do_list_backend/middleware"
	"github.com/lehuynhthang99/learn_to_do_list_backend/models"
)

// GET /to_do_list/:id
func GetAllToDoList(c echo.Context) error {
	customContext, _ := middleware.GetCustomContext(c)

	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return customContext.ErrorWithCustomCode(http.StatusBadRequest, middleware.InvalidInput, err.Error())
	}

	result := new([]models.ToDoList)
	err = config.DB.Where("user_id = ?", userId).Find(&result).Error
	if err != nil {
		return customContext.ErrorWithCustomCode(http.StatusInternalServerError, middleware.DatabaseError, err.Error())
	}

	return customContext.Success(echo.Map{
		"list": result,
	})
}

// POST /to_do_list
func CreateToDoList(c echo.Context) error {
	customContext, _ := middleware.GetCustomContext(c)

	toDoList := new(models.ToDoList)
	err := customContext.Bind(&toDoList)
	if err != nil {
		return customContext.ErrorWithCustomCode(http.StatusBadRequest, middleware.InvalidInput, err.Error())
	}

	result := config.DB.Create(&toDoList)
	if result.Error != nil {
		return customContext.ErrorWithCustomCode(http.StatusInternalServerError, middleware.DatabaseError, result.Error.Error())
	}

	return customContext.Success(toDoList)
}

// PUT /to_do_list/:id/items
func UpdateToDoList(c echo.Context) error {
	customContext, _ := middleware.GetCustomContext(c)

	listIdStr := customContext.Param("id")
	listId, err := strconv.Atoi(listIdStr)
	if err != nil {
		return customContext.ErrorWithCustomCode(http.StatusBadRequest, middleware.InvalidInput, err.Error())
	}

	newItems := new(models.ToDoItems)
	err = customContext.Bind(&newItems)
	if err != nil {
		return customContext.ErrorWithCustomCode(http.StatusBadRequest, middleware.InvalidInput, err.Error())
	}

	toDoList := new(models.ToDoList)
	err = config.DB.Model(&toDoList).Where("id = ?", listId).Update("items", newItems).Error
	if err != nil {
		return customContext.ErrorWithCustomCode(http.StatusInternalServerError, middleware.DatabaseError, err.Error())
	}

	return customContext.Success(nil)
}

// DELETE /to_do_list/:id
func DeleteToDoList(c echo.Context) error {
	customContext, _ := middleware.GetCustomContext(c)
	listIdStr := customContext.Param("id")
	listId, err := strconv.Atoi(listIdStr)
	if err != nil {
		return customContext.ErrorWithCustomCode(http.StatusBadRequest, middleware.InvalidInput, err.Error())
	}

	toDoList := new(models.ToDoList)
	err = config.DB.First(&toDoList, listId).Error
	if err != nil {
		return customContext.ErrorWithCustomCode(http.StatusNotFound, middleware.NotExisted, err.Error())
	}

	err = config.DB.Delete(&toDoList).Error
	if err != nil {
		return customContext.ErrorWithCustomCode(http.StatusInternalServerError, middleware.DatabaseError, err.Error())
	}

	return customContext.Success(nil)
}
