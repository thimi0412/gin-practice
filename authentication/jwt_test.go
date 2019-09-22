package authentication

import (
	"testing"

	"github.com/thimi0412/gin-practice/entity"
)

func TestJWTToken(t *testing.T) {
	user := entity.User{
		ID:       2,
		Email:    "hogehuga@icloud.com",
		Password: "hogehuga",
	}

	tokenString, err := CreateTokenString(user)
	if err != nil {
		t.Fatal(err)
	}

	authUser, err := AuthTokenString(tokenString)
	if err != nil {
		t.Fatal(err)
	}

	if authUser.ID != user.ID {
		t.Fatal("ID不一致")
	}
	if authUser.Email != user.Email {
		t.Fatal("Emain不一致")
	}
	if authUser.Password != authUser.Password {
		t.Fatal("Password不一致")
	}

}
