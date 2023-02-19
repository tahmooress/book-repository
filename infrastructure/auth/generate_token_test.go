package auth

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/tahmooress/book-repository/entities"
)

func TestSuccesssGenerateToken(t *testing.T) {
	const (
		secret   = "my_secret"
		exp      = time.Microsecond
		username = "myUsername"
	)

	authentiactor := Authenticator{
		secretKey: []byte(secret),
		exp:       exp,
	}

	token, err := authentiactor.GenerateToken(username)
	if err != nil {
		t.Error(err)
	}

	err = authentiactor.Verify(token)
	if err != nil {
		t.Error(err)
	}
}

func TestFiledExpGenerateToken(t *testing.T) {
	const (
		secret   = "my_secret"
		exp      = time.Millisecond
		username = "myUsername"
	)

	authentiactor := Authenticator{
		secretKey: []byte(secret),
		exp:       exp,
	}

	token, err := authentiactor.GenerateToken(username)
	if err != nil {
		t.Error(err)
	}

	var e *entities.Err

	err = authentiactor.Verify(token)
	if err != nil && errors.As(err, &e) {
		if !errors.Is(e.Reason, jwt.ErrTokenExpired) {
			t.Error("expiration not work")
		}
	}
}
