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

func ChangeFirstName(c *gin.Context) {
	var firstNameReq firstNameRequest
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}

	customer, err := models.GetCustomer(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Customer not found"})
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
		c.JSON(400, gin.H{"error": "Customer ID is required"})
		return
	}

	customer, err := models.GetCustomer(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Customer not found"})
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
		c.JSON(400, gin.H{"error": "Customer ID is required"})
		return
	}

	customer, err := models.GetCustomer(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Customer not found"})
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
