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

	c.JSON(http.StatusOK, gin.H{"message": "success", "user": u})
}

func RegisterCustomer(c *gin.Context) {
	var customer = models.NewCustomer()
	var dst string
	var wasImageUploaded bool
	var err error

	if err := c.ShouldBind(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with registerInput": err.Error()})
		return
	}

	if err := customer.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"customer.SaveCustomer error": err.Error()})
		return
	}

	if dst, wasImageUploaded, err = customer.HandleUserImageUploads(c, customer.ID, customer.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"image upload error": err.Error()})
		return
	}

	if wasImageUploaded {
		customer.ImagePath = dst
		if err := customer.Update(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"customer.Update error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success!", "customer": customer})
}

func RegisterHost(c *gin.Context) {
	host := models.NewHost()
	var dst string
	var wasImageUploaded bool
	var err error

	if err := c.ShouldBind(&host); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with registerInput": err.Error()})
		return
	}

	if err := host.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"host.SaveHost error": err.Error()})
		return
	}

	if dst, wasImageUploaded, err = host.HandleUserImageUploads(c, host.ID, host.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"image upload error": err.Error()})
		return
	}

	if wasImageUploaded {
		host.ImagePath = dst
		if err := host.Update(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"host.Update error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success!", "host": host})
}

type loginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input loginInput
	var user any

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with loginInput": err.Error()})
		return
	}

	t, user, err := models.LoginCheck(input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"loginCheck: username or password is incorrect": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": t, "user": user})
}
