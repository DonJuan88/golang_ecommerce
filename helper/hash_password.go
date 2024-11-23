package helper

import (
	"github.com/pintarkode/my_api/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	combinedPassword := password + string(config.ENV.TOKEN_LOGIN)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(combinedPassword), bcrypt.DefaultCost)

	return string(passwordHash), err
}
