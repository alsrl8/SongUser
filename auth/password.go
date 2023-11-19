package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPw 텍스트 패스워드를 입력받아 `bcrypt`를 사용하여 hashed 패스워드로 반환한다
func HashPw(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(bytes), err
}

// CheckPwHash 텍스트 패스워드와 hashed 패스워드를 비교하여 동일 여부를 반환한다
func CheckPwHash(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}
