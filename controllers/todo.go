package controllers

import (
	"net/http"
	"todo-app/config"
	"todo-app/models"

	"github.com/labstack/echo/v4"
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

func GetTodos(c echo.Context) (error) {
	var todos []models.Todo
	if result := config.DB.Find(&todos); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error":"error fetching todos"})
	}
	return c.JSON(http.StatusOK, todos)
}

func GetTodoById(c echo.Context) (error) {
	id := c.Param("id")
	var todo models.Todo
	if result := config.DB.First(&todo, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error":"todo not found"})
	}
	return c.JSON(http.StatusOK, todo)
}
