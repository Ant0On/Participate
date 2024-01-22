package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

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
	var customer models.Customer
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

	if dst, wasImageUploaded, err = handleUserImageUploads(c, customer.ID, customer.Role); err != nil {
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

func handleUserImageUploads(c *gin.Context, userID uint, role string) (string, bool, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return "", false, fmt.Errorf("multipart form error: %v", err)
	}

	files := form.File["image"]

	if len(files) > 1 {
		return "", false, fmt.Errorf("only one image can be uploaded, but %d images were provided", len(files))
	}

	if len(files) == 0 {
		fmt.Println("Warning: No image uploaded, using default one instead")
		return "", false, nil
	}

	file := files[0]

	filename := fmt.Sprintf("%d.jpeg", userID)
	var dst string

	if role == "customer" {
		dst = filepath.Join("images/customers", filename)
	} else {
		dst = filepath.Join("images/hosts", filename)
	}

	if err := c.SaveUploadedFile(file, dst); err != nil {
		return "", false, fmt.Errorf("error saving uploaded file: %v", err)
	}

	return dst, true, nil
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

	if dst, wasImageUploaded, err = handleUserImageUploads(c, host.ID, host.Role); err != nil {
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
