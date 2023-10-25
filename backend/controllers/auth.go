package controllers

import (
	"net/http"

	"backend/models"
	"backend/utils/token"

	"github.com/gin-gonic/gin"
)

func CurrentCustomer(c *gin.Context) {
	customerId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"token.ExtractTokenID error": err.Error()})
		return
	}

	u, err := models.GetCustomerById(customerId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"models.GetCustomerByID error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func CurrentHost(c *gin.Context) {
	hostId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"token.ExtractTokenID error": err.Error()})
		return
	}

	u, err := models.GetHostById(hostId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"models.GetHostById error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

// TODO differentiate between Customer and Host register

type registerInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Table     string `json:"table" binding:"required"`
}

func Register(c *gin.Context) {
	var input registerInput
	var customer models.Customer
	var host models.Host

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with registerInput": err.Error()})
		return
	}
	if input.Table == "Customer" {
		customer.Email = input.Email
		customer.Password = input.Password
		if err := customer.Save(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"customer.SaveCustomer error": err.Error()})
			return
		}
	} else if input.Table == "Host" {
		host.Email = input.Email
		host.Password = input.Password
		if err := host.Save(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"host.SaveHost error": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong account type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success!"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error with loginInput": err.Error()})
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
