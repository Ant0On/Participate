package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"backend/models"
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

	if host.Role != "host" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User's role should be host"})
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
		if user.Role == "customer" {
			return fmt.Errorf("invalid role! User is a customer")
		}
		user.Description = value
		return nil
	})
}

func ChangePhoneNumber(c *gin.Context) {
	ChangeField(c, func(user *models.User, value string) error {
		if user.Role == "customer" {
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
		if user.Role == "customer" {
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

func gradeReservation(c *gin.Context, reservationType string) {
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

	var err error
	var reservation interface{}
	switch reservationType {
	case "accommodation":
		var res models.ReservationAccommodation
		err = models.DB.Where("user_id = ? AND ID = ?", customerObj.ID, reservationId).First(&res).Error
		reservation = &res
	case "activity":
		var res models.ReservationActivity
		err = models.DB.Where("user_id = ? AND ID = ?", customerObj.ID, reservationId).First(&res).Error
		reservation = &res
	case "room":
		var res models.ReservationRoom
		err = models.DB.Where("user_id = ? AND ID = ?", customerObj.ID, reservationId).First(&res).Error
		reservation = &res
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	switch res := reservation.(type) {
	case *models.ReservationAccommodation:
		if res.ReservationState != models.Finished {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot grade a reservation that is not finished"})
			return
		}
	case *models.ReservationActivity:
		if res.ReservationState != models.Finished {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot grade a reservation that is not finished"})
			return
		}
	case *models.ReservationRoom:
		if res.ReservationState != models.Finished {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot grade a reservation that is not finished"})
			return
		}
	}

	var request models.Rating
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rate, err := models.GetGradeByCount(strconv.Itoa(request.Count))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch res := reservation.(type) {
	case *models.ReservationAccommodation:
		res.RatingID = rate.ID
		if err := res.Update(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		accommodation, err := models.GetAccommodationById(strconv.Itoa(int(res.AccommodationID)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := accommodation.UpdateRating(request.Count); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	case *models.ReservationActivity:
		res.RatingID = rate.ID
		if err := res.Update(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		activity, err := models.GetActivityById(strconv.Itoa(int(res.ActivityID)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := activity.UpdateRating(request.Count); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	case *models.ReservationRoom:
		res.RatingID = rate.ID
		if err := res.Update(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		room, err := models.GetRoomByID(strconv.Itoa(int(res.RoomID)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		accommodation, err := models.GetAccommodationById(strconv.Itoa(int(room.AccommodationID)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := accommodation.UpdateRating(request.Count); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation graded successfully"})
}

func GradeAccommodationReservation(c *gin.Context) {
	gradeReservation(c, "accommodation")
}

func GradeActivityReservation(c *gin.Context) {
	gradeReservation(c, "activity")
}

func GradeRoomReservation(c *gin.Context) {
	gradeReservation(c, "room")
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

	if user.Role != "customer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only customer can be promoted to host"})
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
