package ecommerce

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)



func CreateProduct(c echo.Context) (error) {
	name := c.FormValue("name")
	description := c.FormValue("description")
	priceStr := c.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error":"invalid input"})
	}
	// Handle the upload
	file, err := c.FormFile("image_url")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error":"image upload error"})
	}
	mimetype
	// open uploaded file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
}
