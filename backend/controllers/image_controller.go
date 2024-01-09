package controllers

import (
	"io"
	"mime/multipart"
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	name := c.PostForm("name")

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{"error": "Image is required"})
		return
	}
	defer file.Close()

	data, err := readFile(header)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read image"})
		return
	}
	// Save image to the database
	err = models.Save(name, data)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read image"})
		return
	}
	c.JSON(200, gin.H{"message": "Image uploaded successfully"})
}

func readFile(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	data, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetImage(c *gin.Context) {
	image, err := models.GetImage()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", image.Data)
}
