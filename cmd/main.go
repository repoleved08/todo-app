package main

import (
	"todo-app/config"
	"todo-app/models"
	"todo-app/routes"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "todo-app/docs"
)

func main() {
	// Initialize database
	config.InitDB()
	// Automigrate the models
	config.DB.AutoMigrate(&models.Todo{}, &models.User{})
	// new echo instance
	e := echo.New()
	// Serve Swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// initialize the routes
	routes.InitRoutes(e)
	// start server
	e.Logger.Fatal(e.Start(":8080"))
}
