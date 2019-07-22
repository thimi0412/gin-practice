package authentication

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/thimi0412/gin-practice/db"
	"github.com/thimi0412/gin-practice/entity"
)

// JWTToken jwt作成用
type JWTToken struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"e_mail"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// CreateTokenString 承認用jwtを作成
func CreateTokenString(user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &JWTToken{
		UserID:   user.ID,
		Email:    user.Email,
		Password: user.Password,
	})
	tokenString, err := token.SignedString([]byte("foobar"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AuthTokenString jwtからユーザー情報を取得する
func AuthTokenString(tokenString string) (entity.User, error) {
	user := entity.User{}
	jwtToken := JWTToken{}
	_, err := jwt.ParseWithClaims(tokenString, &jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	})
	if err != nil {
		return user, err
	}

	conn := db.DBConnect()
	defer conn.Close()

	user.ID = jwtToken.UserID

	if err := conn.First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
