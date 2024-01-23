package token

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email, role string) (string, error) {
	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", fmt.Errorf("strconv.Atoi: %w", err)
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	if email == "admin@participate.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = role
	}
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func IsTokenValid(c *gin.Context, role string) error {
	token, err := getToken(c)
	if err != nil {
		return fmt.Errorf("IsTokenValid: %w", err)
	}
	extractedRole, err := extractRole(token)
	if extractedRole != role {
		return fmt.Errorf("unauthorized role - expected: %s, got: %s", role, extractedRole)
	}
	return nil
}

func getToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("jwt.Parse: %w", err)
	}
	return token, nil
}

func extractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func extractRole(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uRole := fmt.Sprintf("%s", claims["role"])
		return uRole, nil
	}
	return "", nil
}

func ExtractTokenEmail(c *gin.Context) (string, error) {
	token, err := getToken(c)
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uEmail := fmt.Sprintf("%s", claims["email"])
		if err != nil {
			return "", fmt.Errorf("strconv.ParseUint: %w", err)
		}
		return uEmail, nil
	}
	return "", nil
}
