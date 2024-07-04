package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckState(state string) bool {
	switch state {
	case "pending":
		return true
	case "accepted":
		return true
	case "ongoing":
		return true
	case "finished":
		return true
	case "rejected":
		return true
	default:
		return false
	}
}

type ChangePriceReq struct {
	Price float64 `json:"price" binding:"required,gt=1"`
}

func HistoryWhereCondition(userID string) string {
	return fmt.Sprintf("app_user.id = '%s' AND reservation_state in ('finished', 'pending', 'accepted', 'rejected')", userID)
}

func CurrentWhereCondition(userID string) string {
	return fmt.Sprintf("app_user.id = '%s' AND reservation_state in ('pending', 'accepted')", userID)
}

func CheckToken(c *gin.Context) {
	c.String(http.StatusOK, "Token is valid")
}

type DiscountRequest struct {
	Discount float64 `json:"discount" binding:"required,gte=0,lte=100"`
}

type EquipmentRequest struct {
	Equipment []string `json:"equipment"`
}

type FacilitiesRequest struct {
	Facilities []string `json:"facilities"`
}
