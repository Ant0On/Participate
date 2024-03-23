package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"backend/models"
	"backend/models/DTO"
	"backend/pkg/passHelper"

	"github.com/gin-gonic/gin"
)

type fieldChangeRequest struct {
	Value string `json:"value" binding:"required"`
}

type promoteRequest struct {
	Password    string `json:"password" binding:"required,min=8"`
	Description string `json:"description" binding:"required,min=15,max=255"`
	PhoneNumber string `json:"phone_number" binding:"required,numeric,min=9,max=15"`
	BankAccount string `json:"bank_account" binding:"required,numeric,min=16,max=40"`
}

type passwordRequest struct {
	OldPassword     string `json:"old_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=8,nefield=OldPassword"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword"`
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

func getFieldChangeResponse(c *gin.Context, id string, updateFunc func(*models.User, string) error) {
	var req fieldChangeRequest

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := updateFunc(user, req.Value); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"user.Update": err.Error()})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"message": "success", "user": user})
}

func ChangeField(c *gin.Context, updateFunc func(*models.User, string) error) {
	id := c.Param("id")
	getFieldChangeResponse(c, id, updateFunc)
}

func ChangeFirstName(c *gin.Context) {
	ChangeField(c, func(user *models.User, value string) error {
		if len(value) < 2 || len(value) > 100 {
			return fmt.Errorf("first name must be between 2 and 100 characters")
		}
		user.FirstName = value
		return nil
	})
}

func ChangeLastName(c *gin.Context) {
	ChangeField(c, func(user *models.User, value string) error {
		if len(value) < 2 || len(value) > 100 {
			return fmt.Errorf("last name must be between 2 and 100 characters")
		}
		user.LastName = value
		return nil
	})
}

func ChangeEmail(c *gin.Context) {
	ChangeField(c, func(user *models.User, value string) error {
		if !isValidEmail(value) {
			return fmt.Errorf("invalid email address")
		}
		user.Email = value
		return nil
	})
}

func ChangeDescription(c *gin.Context) {
	ChangeField(c, func(user *models.User, value string) error {
		if user.Role != "customer" {
			return fmt.Errorf("invalid role! User is a customer")
		}
		user.Description = value
		return nil
	})
}

func ChangePhoneNumber(c *gin.Context) {
	ChangeField(c, func(user *models.User, value string) error {
		if user.Role != "customer" {
			return fmt.Errorf("invalid role! User is a customer")
		}
		if len(value) < 9 || len(value) > 15 {
			return fmt.Errorf("phone number must be between 9 and 15 digits")
		}
		user.PhoneNumber = value
		return nil
	})
}

func ChangeBankAccount(c *gin.Context) {
	ChangeField(c, func(user *models.User, value string) error {
		if user.Role != "customer" {
			return fmt.Errorf("invalid role! User is a customer")
		}
		if len(value) < 16 || len(value) > 40 {
			return fmt.Errorf("bank account must be between 16 and 40 digits")
		}
		user.BankAccount = value
		return nil
	})
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func ChangeImage(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
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

func ChangePassword(c *gin.Context) {
	var passwordReq passwordRequest
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&passwordReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := passHelper.VerifyPassword(passwordReq.OldPassword, user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Old password is incorrect!"})
		return
	}

	user.Password = passwordReq.NewPassword

	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"user.HashPassword": err.Error()})
		return
	}

	if err := user.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"user.Update": err.Error()})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func GradeAccommodationReservation(c *gin.Context) {
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

	customerObj, ok := customer.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var reservation models.ReservationAccommodation
	err := models.DB.Where("user_id = ? AND ID = ?", customerObj.ID, reservationId).First(&reservation).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	if reservation.ReservationState != models.Finished {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot grade a reservation that is not finished"})
		return
	}

	var request models.Rating
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rate, err := models.GetGradeByCount(strconv.Itoa(request.Count))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	reservation.GradeID = rate.ID

	if err := reservation.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation graded successfully"})
}

func GradeActivityReservation(c *gin.Context) {
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

	customerObj, ok := customer.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var reservation models.ReservationActivity
	err := models.DB.Where("user_id = ? AND ID = ?", customerObj.ID, reservationId).First(&reservation).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	if reservation.ReservationState != models.Finished {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot grade a reservation that is not finished"})
		return
	}

	var request models.Rating
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rate, err := models.GetGradeByCount(strconv.Itoa(request.Count))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	reservation.RatingID = rate.ID

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
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
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
	c.JSON(http.StatusOK, gin.H{"message": "User upgraded to Host successfully", "host": user})
}

func GetReservationsAccommodationHistory(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	var finishedReservations []DTO.ReservationAccommodationWithOffer
	result := models.DB.
		Model(&models.ReservationAccommodation{}).
		Joins("JOIN accommodation ON reservation_accommodation.offer_id = accommodation.id").
		Joins("JOIN town ON accommodation.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Joins("JOIN app_user ON reservation_accommodation.user_id = app_user.id").
		Where("app_user.id = ? AND reservation_state in ('finished', 'accepted', 'rejected')", userID).
		Select("reservation_accommodation.id as reservation_id, reservation_accommodation.date_from, reservation_accommodation.date_to, reservation_accommodation.capacity," +
			" reservation_accommodation.rating_id, accommodation.title, accommodation.price_per_day, accommodation.is_animal_friendly, accommodation.accommodation_type, town.name as town_name, country.name as country_name, reservation_accommodation.reservation_state, accommodation.id as offer_id").
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
		Joins("JOIN app_user ON offer.user_id = app_user.id").
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
