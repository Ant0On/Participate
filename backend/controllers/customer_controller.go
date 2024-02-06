package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

type firstNameRequest struct {
	FirstName string `json:"first_name"`
}
type lastNameRequest struct {
	LastName string `json:"last_name"`
}
type emailRequest struct {
	Email string `json:"email"`
}

type promoteRequest struct {
	Description string `json:"description" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	BankAccount string `json:"bank_account" binding:"required"`
}

func ChangeFirstName(c *gin.Context) {
	var firstNameReq firstNameRequest
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	customer, err := models.GetCustomer(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if err := c.ShouldBindJSON(&firstNameReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.FirstName = firstNameReq.FirstName

	if err := customer.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"customer.Update": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "customer": customer})
}
func ChangeLastName(c *gin.Context) {
	var lastNameReq lastNameRequest
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer ID is required"})
		return
	}

	customer, err := models.GetCustomer(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if err := c.ShouldBindJSON(&lastNameReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.LastName = lastNameReq.LastName

	if err := customer.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"customer.Update": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "customer": customer})
}
func ChangeEmail(c *gin.Context) {
	var emailReq emailRequest
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer ID is required"})
		return
	}

	customer, err := models.GetCustomer(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if err := c.ShouldBindJSON(&emailReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.Email = emailReq.Email

	if err := customer.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"customer.Update": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "customer": customer})
}

func PromoteToHost(c *gin.Context) {
	id := c.Param("id")
	var promoteReq promoteRequest

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer ID is required"})
		return
	}

	customer, err := models.GetCustomer(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if err := c.ShouldBindJSON(&promoteReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	host := models.Host{
		Customer:    customer,
		Description: promoteReq.Description,
		PhoneNumber: promoteReq.PhoneNumber,
		BankAccount: promoteReq.BankAccount,
		Offers:      nil,
	}
	if err := host.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := customer.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer upgraded to Host successfully", "host": host})
}
