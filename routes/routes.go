package routes

import (
	"todo-app/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/api/todo/create", controllers.CreateTodo)
	e.GET("/api/todos", controllers.GetTodos)
	e.GET("/api/todo/:id", controllers.GetTodoById)
}
