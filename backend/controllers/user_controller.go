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
	FirstName string `json:"first_name" binding:"required"`
}

type lastNameRequest struct {
	LastName string `json:"last_name" binding:"required,min=2,max=100"`
}

type emailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type promoteRequest struct {
	Password    string `json:"password" binding:"required,min=8"`
	Description string `json:"description" binding:"required,min=15,max=255"`
	PhoneNumber string `json:"phone_number" binding:"required,numeric,min=9,max=15"`
	BankAccount string `json:"bank_account" binding:"required,numeric,min=16,max=40"`
}

func GetHostByID(c *gin.Context) {
	hostID := c.Param("id")

	if hostID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Host ID is required"})
		return
	}

	host, err := models.GetUserById(hostID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Host not found"})
		return
	}

	host.Password = ""

	c.JSON(http.StatusOK, host)
}

func ChangeFirstName(c *gin.Context) {
	var firstNameReq firstNameRequest
	id := c.Param("id")

	user, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AppUser not found"})
		return
	}

	if err := c.ShouldBindJSON(&firstNameReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.FirstName = firstNameReq.FirstName

	if err := user.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"user.Update": err.Error()})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"message": "success", "user": user})
}
func ChangeLastName(c *gin.Context) {
	var lastNameReq lastNameRequest
	id := c.Param("id")

	user, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AppUser not found"})
		return
	}

	if err := c.ShouldBindJSON(&lastNameReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.LastName = lastNameReq.LastName

	if err := user.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"user.Update": err.Error()})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"message": "success", "user": user})
}
func ChangeEmail(c *gin.Context) {
	var emailReq emailRequest
	id := c.Param("id")

	user, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AppUser not found"})
		return
	}

	if err := c.ShouldBindJSON(&emailReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Email = emailReq.Email

	if err := user.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"user.Update": err.Error()})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"message": "success", "user": user})
}

func ChangeImage(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AppUser ID is required"})
		return
	}

	user, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AppUser not found"})
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

	if err := user.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"user.Update": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "user": user})
}

func GradeReservation(c *gin.Context) {
	reservationId := c.Param("id")
	if reservationId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}

	customer, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	customerObj, ok := customer.(*models.AppUser)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var reservation models.Reservation
	err := models.DB.Where("app_user_id = ? AND ID = ?", customerObj.ID, reservationId).First(&reservation).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	if reservation.ReservationState != models.Finished {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot grade a reservation that is not finished"})
		return
	}

	var request models.Grade
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

	user, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AppUser not found"})
		return
	}

	if err := c.ShouldBindJSON(&promoteReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := passHelper.VerifyPassword(promoteReq.Password, user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Passwords do not match"})
		return
	}

	user.Role = "host"
	user.Description = promoteReq.Description
	user.PhoneNumber = promoteReq.PhoneNumber
	user.BankAccount = promoteReq.BankAccount

	if err := user.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"message": "AppUser upgraded to Host successfully", "host": user})
}

func GetReservationsHistory(c *gin.Context) {
	userID := c.Param("id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AppUser ID is required"})
		return
	}

	var finishedReservations []DTO.ReservationWithOffer
	result := models.DB.
		Model(&models.Reservation{}).
		Joins("JOIN offer ON reservation.offer_id = offer.id").
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Joins("JOIN app_user ON reservation.app_user_id = app_user.id").
		Where("app_user.id = ? AND reservation_state in ('finished', 'accepted', 'rejected')", userID).
		Select("reservation.id as reservation_id, reservation.date_from, reservation.date_to, reservation.number_of_people, reservation.grade_id, offer.name," + "" +
			"offer.price, offer.is_animal_friendly, offer.offer_type, town.name as town_name, country.name as country_name, reservation.reservation_state, offer.id as offer_id").
		Find(&finishedReservations)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNoContent, gin.H{"warning": "No finished reservations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "finished reservations fetched successfully", "data": finishedReservations})
}

func GetPendingReservations(c *gin.Context) {
	userID := c.Param("id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "app_user ID is required"})
		return
	}

	var pendingReservations []DTO.ReservationWithOffer
	result := models.DB.
		Model(&models.Reservation{}).
		Joins("JOIN offer ON reservation.offer_id = offer.id").
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Joins("JOIN app_user ON offer.app_user_id = app_user.id").
		Where("app_user.id = ? AND reservation_state = 'pending'", userID).
		Select("reservation.id as reservation_id, reservation.date_from, reservation.date_to, reservation.number_of_people, offer.name," + "" +
			"offer.price, offer.is_animal_friendly, offer.offer_type, town.name as town_name, country.name as country_name, offer.id as offer_id").
		Find(&pendingReservations)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNoContent, gin.H{"error": "No pending reservations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pending reservations fetched successfully", "data": pendingReservations})
}
