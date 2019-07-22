package authentication

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// PasswordHash パスワード暗号化する
func PasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hash), err
}

// PasswordVerify パスワードがあっているか確認する
func PasswordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
