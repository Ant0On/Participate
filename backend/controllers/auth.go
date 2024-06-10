package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

type registerInput struct {
	FirstName            string `form:"first_name" binding:"required,min=2,max=30"`
	LastName             string `form:"last_name" binding:"required,min=2,max=30"`
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"required,min=8"`
	PasswordConfirmation string `form:"password_confirmation" binding:"required,eqfield=Password"`
}

func Register(c *gin.Context) {
	var user models.User
	var input registerInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Validation failed: " + err.Error()})
		return
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Password = input.Password

	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "Failed to hash user password: " + err.Error()})
		return
	}

	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "Failed to save user: " + err.Error()})
		return
	}

	dst, wasImageUploaded, err := user.HandleUserImageUploads(c, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Failed to upload user image: " + err.Error()})
		return
	}

	if wasImageUploaded {
		user.ImagePath = dst
		if err := user.Update(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "Failed to update user image path: " + err.Error()})
			return
		}
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"message": "registration success!", "user": user})
}

type loginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input loginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Failed to login: " + err.Error()})
		return
	}

	t, user, err := models.LoginCheck(input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Failed to login: " + err.Error()})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"token": t, "user": user})
}
