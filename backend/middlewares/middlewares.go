package middlewares

import (
	"net/http"

	"backend/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.IsTokenValid(c, role)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
