package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"todo-app/config"
	"todo-app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateTodo(c echo.Context) error {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}
	if result := config.DB.Create(&todo); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create todo"})
	}
	return c.JSON(http.StatusCreated, todo)
}

func GetTodos(c echo.Context) error {
	var todos []models.Todo
	if result := config.DB.Find(&todos); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "error fetching todos"})
	}
	return c.JSON(http.StatusOK, todos)
}

// @Summary Get a Todo
// @Description Get details of a Todo
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/todos/{id} [get]
func GetTodoById(c echo.Context) error {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid todo id"})
    }

    var todo models.Todo
    if result := config.DB.First(&todo, id); result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return c.JSON(http.StatusNotFound, echo.Map{"error": "todo not found"})
        }
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "internal server error"})
    }

    return c.JSON(http.StatusOK, todo)
}


func UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	var todo models.Todo
	if result := config.DB.First(&todo, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "todo not found"})
	}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}
	config.DB.Save(&todo)
	return c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	if result := config.DB.Delete(&models.Todo{}, id); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to delete todo"})
	}
	return c.NoContent(http.StatusNoContent)
}
