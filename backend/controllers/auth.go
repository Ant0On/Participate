package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"backend/models"
	"backend/utils/token"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var customer models.Customer

	if err := c.ShouldBind(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with registerInput": err.Error()})
		return
	}

	customer.ID = uuid.New().String()

	if err := handleUserImageUploads(c, customer.ID, customer.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"image upload error": err.Error()})
		return
	}

	if err := customer.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"customer.SaveCustomer error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration success!", "customer": customer})
}

func handleUserImageUploads(c *gin.Context, userID, role string) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["images"]
	for i, file := range files {
		filename := fmt.Sprintf("%s_%d.jpeg", userID, i)
		dst := filepath.Join("images/offers", userID, filename)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			return err
		}
	}

	return nil
}

func RegisterHost(c *gin.Context) {
	host := models.NewHost()

	if err := c.ShouldBindJSON(&host); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with registerInput": err.Error()})
		return
	}

	if err := host.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"host.SaveHost error": err.Error()})
		return
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
