package controllers

import (
	"net/http"

	"backend/models"
	"backend/utils/token"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	userEmail, err := token.ExtractTokenEmail(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"token.ExtractTokenEmail error": err.Error()})
		return
	}

	u, err := models.GetUserByEmail(userEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"models.GetUserByEmail error": err.Error()})
		return
	}

	u.Password = ""

	c.JSON(http.StatusOK, gin.H{"message": "success", "user": u})
}

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
		c.JSON(http.StatusBadRequest, gin.H{"error with registerInput": err.Error()})
		return
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Password = input.Password

	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user.HashPassword error": err.Error()})
		return
	}

	if err := user.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user.SaveCustomer error": err.Error()})
		return
	}

	dst, wasImageUploaded, err := user.HandleUserImageUploads(c, user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"image upload error": err.Error()})
		return
	}

	if wasImageUploaded {
		user.ImagePath = dst
		if err := user.Update(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"user.Update error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error with loginInput": err.Error()})
		return
	}

	t, user, err := models.LoginCheck(input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"loginCheck: username or password is incorrect": err.Error()})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"token": t, "user": user})
}
