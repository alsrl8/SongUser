package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 텍스트 패스워드를 입력받아 `bcrypt`를 사용하여 hashed 패스워드로 반환한다
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash 텍스트 패스워드와 hashed 패스워드를 비교하여 동일 여부를 반환한다
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
