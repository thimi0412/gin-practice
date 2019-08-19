package authentication

import "testing"

func TestCrypt(t *testing.T) {
	password := "hogehuga"
	hash, err := PasswordHash(password)
	if err != nil {
		t.Fatal(err)
	}

	if err := PasswordVerify(hash, password); err != nil {
		t.Fatal(err)
	}
}
