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

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

type registerCustomerInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func RegisterCustomer(c *gin.Context) {
	var input registerCustomerInput
	var customer *models.Customer

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with registerInput": err.Error()})
		return
	}

	customer = assignCustomerInput(&input)

	if err := customer.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"customer.SaveCustomer error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration success!"})
}

func assignCustomerInput(input *registerCustomerInput) *models.Customer {
	var customer models.Customer
	customer.FirstName = input.FirstName
	customer.LastName = input.LastName
	customer.Email = input.Email
	customer.Password = input.Password

	return &customer
}

type registerHostInput struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `gorm:"size:12;not null;unique" json:"phone_number" binding:"required"`
	BankAccount string `gorm:"size:31;not null;unique" json:"bank_account" binding:"required"`
}

func RegisterHost(c *gin.Context) {
	var input registerHostInput
	var host *models.Host

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with registerInput": err.Error()})
		return
	}

	host = assignHostInput(input)

	if err := host.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"host.SaveHost error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success!"})
}

func assignHostInput(input registerHostInput) *models.Host {
	host := models.NewHost()
	host.FirstName = input.FirstName
	host.LastName = input.LastName
	host.Email = input.Email
	host.Password = input.Password
	host.PhoneNumber = input.PhoneNumber
	host.BankAccount = input.BankAccount

	return host
}

type loginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Table    string `json:"table" binding:"required"`
}

func Login(c *gin.Context) {
	var input loginInput
	var user models.Customer

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with loginInput": err.Error()})
		return
	}

	user.Email = input.Email
	user.Password = input.Password
	t, err := user.LoginCheck(user.Email, user.Password, input.Table)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"loginCheck: username or password is incorrect": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": t})
}
