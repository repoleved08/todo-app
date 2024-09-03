package controllers

import (
	"net/http"
	"os"
	"time"
	"todo-app/config"
	"todo-app/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Register a new user
// @Description Create a new user account
// @Produce json
// @Param user body models.User true "User Registration"
// @Success 201 {object} models.User
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/auth/register [post]
func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input", "details": err.Error()})
	}

	// Validate user input (you might want to add more validation here)
	if user.Username == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Username and password are required"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error hashing password", "details": err.Error()})
	}
	user.Password = string(hashedPassword)

	// Save the user to the database
	if result := config.DB.Create(&user); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Account creation failed", "details": result.Error.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

// @Summary Login user
// @Description Authenticate user and return a JWT token
// @Produce json
// @Param user body models.User true "User Login"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Router /api/auth/login [post]
func Login(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input", "details": err.Error()})
	}

	var dbUser models.User
	if err := config.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid username or password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  dbUser.ID,
		"username": dbUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error generating token", "details": err.Error()})
	}

	response := models.LoginResponse{
		User: models.UserDto{
			Username: dbUser.Username,
			Email:    dbUser.Email,
		},
		Token: tokenString,
	}

	return c.JSON(http.StatusOK, response)
}
