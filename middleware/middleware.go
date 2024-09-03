package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Authorization header required"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			fmt.Println("Error Parsing Token: ", err)
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
		}

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token claims"})
		}

		if exp, ok := claims["exp"].(float64); !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Token has expired"})
		}

		c.Set("user_id", claims["user_id"])
		return next(c)
	}
}
