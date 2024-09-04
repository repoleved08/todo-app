package controllers

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"todo-app/config"
	"todo-app/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	priceStr := c.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}
	// Handle the upload
	file, err := c.FormFile("image_url")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "image upload error"})
	}
	mimetype := file.Header.Get("Content-Type")
	if mimetype != "image/jpeg" && mimetype != "image/png" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid image format"})
	}
	// open uploaded file
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "error opening the file"})
	}
	defer src.Close()
	// check if the file is avalid image by decoding it
	_, _, err = image.Decode(src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid image file"})
	}
	// reset file pointer to the beginning
	if _, err := src.Seek(0, 0); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "error resetting the file pointer"})
	}
	uniqueFileName := uuid.New().String() + filepath.Ext(file.Filename)

	// Create a destination file on the server
	uploadDir := "static/products/" // Make sure this directory exists and is writeable
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return err
	}

	// Construct the full file path
	filePath := filepath.Join(uploadDir, uniqueFileName)

	// Create the destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the uploaded file data to the destination file
	if _, err = src.Seek(0, 0); err != nil {
		return err
	}
	if _, err = dst.ReadFrom(src); err != nil {
		return err
	}

	// Create an accessible URL for the image
	// Assuming your NGINX is serving files from the "uploads" directory under "/static/" path
	domain := os.Getenv("APP_DOMAIN")
	imageURL := fmt.Sprintf("http://%s/static/products/%s", domain, uniqueFileName)

	// Get the user ID from the context
	userID := uint(c.Get("user_id").(float64))

	// Create the product
	product := models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		ImageURL:    imageURL, // Save image URL instead of file path
		UserID:      userID,
	}

	// Save product to the database
	config.DB.Create(&product)
	return c.JSON(http.StatusCreated, product)
}

func GetProducts(c echo.Context) (error) {
	var products []models.Product
	if result := config.DB.Find(&products); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error":"error fetching products"})
	}
	return c.JSON(http.StatusOK, products)
}
