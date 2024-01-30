package middlewares

import (
	"net/http"

	"backend/models"
	"backend/utils/token"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		getToken, err := token.GetToken(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		claims, ok := getToken.Claims.(jwt.MapClaims)
		if !ok || !getToken.Valid {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		userID := claims["user_id"].(uint)
		if err := token.IsTokenValid(c, role); err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized, wrong role")
			c.Abort()
			return
		}

		var cust models.Customer
		var h models.Host

		if role == "customer" {
			if err := models.DB.First(&cust, userID).Error; err != nil {
				c.String(http.StatusUnauthorized, "User not found")
				c.Abort()
				return
			}
			c.Set("user", &cust)
		} else {
			if err := models.DB.First(&h, userID).Error; err != nil {
				c.String(http.StatusUnauthorized, "User not found")
				c.Abort()
				return
			}
			c.Set("user", &h)
		}
		c.Set("role", role)
		c.Next()
	}
}
