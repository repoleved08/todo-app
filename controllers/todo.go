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

// @Summary Create a new Todo
// @Description Create a new Todo
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Todo data"
// @Success 201 {object} models.Todo
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/todos/create [post]
// @Security BearerToken
func CreateTodo(c echo.Context) error {
	var todo models.Todo

	// Bind the incoming request to the todo object
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	// Additional validation (if needed)
	if todo.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "title is required"})
	}

	// Create the new todo in the database
	if result := config.DB.Create(&todo); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create todo", "details": result.Error.Error()})
	}

	// Return the created todo with HTTP 201 status
	return c.JSON(http.StatusCreated, todo)
}

// @Summary Get all Todos
// @Description Get all Todos
// @Produce json
// @Success 200 {object} []models.Todo
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/todos [get]
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

// @Summary Update a Todo
// @Description Update a Todo by ID
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body models.Todo true "Todo data"
// @Success 200 {object} models.Todo
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/todos/update/{id} [put]
// @Security BearerToken
func UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	var todo models.Todo
	if result := config.DB.First(&todo, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "todo not found"})
	}

	// Bind the incoming JSON request to the todo object
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	// Save the updated todo
	config.DB.Save(&todo)
	return c.JSON(http.StatusOK, todo)
}

// @Summary Delete a Todo
// @Description Delete a Todo by ID
// @Produce json
// @Param id path int true "Todo ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /api/todos/delete/{id} [delete]
// @Security BearerToken
func DeleteTodo(c echo.Context) error {
	id := c.Param("id")

	// Check if the Todo exists
	var todo models.Todo
	if result := config.DB.First(&todo, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "todo not found"})
	}

	// Delete the Todo
	if result := config.DB.Delete(&todo); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to delete todo", "details": result.Error.Error()})
	}

	// Return success message
	return c.NoContent(http.StatusNoContent)
}
