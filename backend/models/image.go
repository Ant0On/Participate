package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
}

func CreateImage(c *gin.Context) {
	var image Image

	// Bind the form data to the image struct
	if err := c.ShouldBind(&image); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(200, gin.H{"message": "User created successfully", "image": image})
}
