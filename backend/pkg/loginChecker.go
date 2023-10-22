package pkg

import "github.com/gin-gonic/gin"

type LoginChecker interface {
	LoginCheck(email, password string) (string, error)
	AccountType(ctx *gin.Context) (string, error)
}
