package usecases

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "somepassword@"

	hash, err := hashPassword(password)
	if err != nil {
		t.Error(err)
	}

	ok, err := comparePasswords(hash, password)
	if err != nil {
		t.Error(err)
	}

	if !ok {
		t.Errorf("compare failed")
	}
}
