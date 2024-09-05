package routes

import (
	"todo-app/controllers"
	"todo-app/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	// Authentication
	e.POST("/api/auth/register", controllers.Register)
	e.POST("/api/auth/login", controllers.Login)
	// Todo
	e.GET("/api/todos", controllers.GetTodos)
	e.GET("/api/todos/:id", controllers.GetTodoById)
	// protected routes
	e.POST("/api/todos/create", controllers.CreateTodo, middleware.JWTMiddleware)
	e.PUT("/api/todos/update/:id", controllers.UpdateTodo, middleware.JWTMiddleware)
	e.DELETE("/api/todos/delete/:id", controllers.DeleteTodo, middleware.JWTMiddleware)
	// e-commerce endpoints
	e.POST("/api/products/create", controllers.CreateProduct, middleware.JWTMiddleware)
	e.GET("/api/products", controllers.GetProducts)
	e.GET("/api/products/:id", controllers.GetProductById)
	e.PUT("/api/products/update/:id", controllers.UpdateProduct, middleware.JWTMiddleware)
	e.DELETE("/api/products/delete/:id", controllers.DeleteProduct, middleware.JWTMiddleware)
	
}
