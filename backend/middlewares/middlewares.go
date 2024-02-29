package middlewares

import (
	"net/http"

	"backend/models"
	"backend/utils/token"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthMiddleware(requiredRole string) gin.HandlerFunc {
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

		userID := claims["user_id"].(float64)
		userRole := claims["role"].(string)

		if userRole != requiredRole && requiredRole != "customer" {
			c.String(http.StatusUnauthorized, "Unauthorized, wrong role")
			c.Abort()
			return
		}

		if requiredRole != "customer" && c.FullPath() == "/api/customer/:id/promote" {
			c.String(http.StatusUnauthorized, "Unauthorized for promotion, wrong role")
			c.Abort()
			return
		}

		var user models.User

		if err := models.DB.First(&user, int(userID)).Error; err != nil {
			c.String(http.StatusUnauthorized, "User not found")
			c.Abort()
			return
		}
		c.Set("user", &user)
		c.Set("role", userRole)
		c.Next()
	}
}
