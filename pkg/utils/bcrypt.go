package utils

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var bcryptCost = 16

func HashPassword(password string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		logrus.Errorf("[Bcrypt Password] error while generate from password, err: %v", err)
		return ""
	}
	return string(b)
}

func CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
