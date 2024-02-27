package passHelper

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password, hashedPassword string) error {
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)))
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
