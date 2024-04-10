package utils

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

type DiscountRequest struct {
	Discount float64 `json:"discount" binding:"required,gte=0,lte=100"`
}
