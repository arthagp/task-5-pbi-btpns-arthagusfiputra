package hash //the name of module in this file , in js like a module.export?/ just maybe

import (
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

// compare password 
func CheckPasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}