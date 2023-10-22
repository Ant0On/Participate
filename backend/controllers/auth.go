package controllers

import (
	"fmt"
	"net/http"

	"backend/models"
	"backend/utils/token"

	"github.com/gin-gonic/gin"
)

func CurrentCustomer(c *gin.Context) {
	customerId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("token.ExtractTokenID: %w", err)})
		return
	}

	u, err := models.GetCustomerByID(customerId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("models.GetCustomerByID: %w", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

type loginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Table    string `json:"table" binding:"required"`
}

func Login(c *gin.Context) {
	var input loginInput
	var customer models.Customer
	var host models.Host

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("error with loginInput: %w", err)})
		return
	}

	if input.Table == "Customer" {
		customer.Email = input.Email
		customer.Password = input.Password
		t, err := customer.AccountType(c)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": t})
	} else {
		host.Email = input.Email
		host.Password = input.Password
		t, err := host.AccountType(c)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": t})
	}

}

// TODO differentiate between Customer and Host register

type registerInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input registerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("error with registerInput: %w", err)})
		return
	}

	customer := models.Customer{}

	customer.Email = input.Email
	customer.Password = input.Password

	err := customer.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("customer.SaveCustomer: %w", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success!"})
}
