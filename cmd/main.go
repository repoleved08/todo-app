package main

import (
	"todo-app/config"
	"todo-app/models"
	"todo-app/routes"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "todo-app/docs"
)

// @title Swagger Techxtrasol API
// @version 1.0
// @description This is a server for managing everything.
// @securityDefinitions.bearer BearerToken
// @in header
// @name Authorization
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080


func main() {
	// Initialize database
	config.InitDB()
	// Automigrate the models
	config.DB.AutoMigrate(&models.Todo{}, &models.User{}, &models.Product{})
	// new echo instance
	e := echo.New()
	// Serve Swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// initialize the routes
	routes.InitRoutes(e)
	// start server
	e.Logger.Fatal(e.Start(":8080"))
}
