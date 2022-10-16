package crypto

import (
	"golang.org/x/crypto/bcrypt"
)
var (
	CryptoGenerate = bcrypt.GenerateFromPassword
	CryptoCompare = bcrypt.CompareHashAndPassword
)
func HashPassword(password string) (string, error) {
	bytes, err := CryptoGenerate([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := CryptoCompare([]byte(hash), []byte(password))
	return err == nil
}
