package main

import (
	"todo-app/config"
	"todo-app/models"
	"todo-app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize database
	config.InitDB()
	// Automigrate the models
	config.DB.AutoMigrate(&models.Todo{})
	// new echo instance
	e := echo.New()
	// initialize the routes
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
