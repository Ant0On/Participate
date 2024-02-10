package controllers

import (
	"net/http"

	"strconv"

	"backend/models"
	"backend/models/DTO"
	"backend/pkg/passHelper"

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
	Password    string `json:"password" binding:"required"`
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

func GradeReservation(c *gin.Context) {
	offerID := c.Param("id")
	if offerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offer ID"})
		return
	}

	customer, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	customerObj, ok := customer.(*models.Customer)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var reservation models.Reservation
	err := models.DB.Where("customer_id = ? AND offer_id = ?", customerObj.ID, offerID).First(&reservation).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	if reservation.ReservationState != models.Finished {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot grade a reservation that is not finished"})
		return
	}

	var request *models.Grade
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grade, err := models.GetGradeByCount(strconv.Itoa(request.Count))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	reservation.GradeID = grade.ID

	if err := reservation.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation graded successfully"})
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

	if err := passHelper.VerifyPassword(promoteReq.Password, customer.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Passwords do not match"})
		return
	}

	customer.Role = "host"
	customer.Password = promoteReq.Password

	host := models.Host{
		Customer:    customer,
		Description: promoteReq.Description,
		PhoneNumber: promoteReq.PhoneNumber,
		BankAccount: promoteReq.BankAccount,
		Offers:      nil,
	}

	if err := customer.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := host.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer upgraded to Host successfully", "host": host})
}

func GetReservationsHistory(c *gin.Context) {
	customerID := c.Param("id")

	if customerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer ID is required"})
		return
	}

	var finishedReservations []DTO.ReservationWithOffer
	result := models.DB.
		Model(&models.Reservation{}).
		Joins("JOIN offer ON reservation.offer_id = offer.id").
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Joins("JOIN customer ON reservation.customer_id = customer.id").
		Where("customer.id = ? AND reservation_state in ('finished', 'accepted', 'rejected')", customerID).
		Select("reservation.id as reservation_id, reservation.date_from, reservation.date_to, offer.name," + "" +
			"offer.price, offer.is_animal_friendly, offer.offer_type, town.name as town_name, country.name as country_name, reservation.reservation_state, offer.id as offer_id").
		Find(&finishedReservations)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No finished reservations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "finished reservations fetched successfully", "data": finishedReservations})
}
