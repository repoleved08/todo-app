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


func Register(c echo.Context) error {
	var user = new(models.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "error hashing password"})
	}
	user.Password = string(hashedPassword)
	if result := config.DB.Create(&user); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Account creation failed!"})
	}
	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var dbUser models.User
	if err := config.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid username or password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  dbUser.ID,
		"username": dbUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "error generating token"})
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

