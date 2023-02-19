package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/tahmooress/book-repository/entities"
	"github.com/tahmooress/book-repository/pkg/util"
)

var ErrMisMatchAlg = errors.New("signing algorithm is not match")

const defaultExp = time.Minute * 5

type Claims struct {
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

type Authenticator struct {
	secretKey []byte
	exp       time.Duration
}

func New(secretKey []byte, exp string) *Authenticator {
	return &Authenticator{
		secretKey: secretKey,
		exp:       util.ParseDurationWithDefault(exp, defaultExp),
	}
}

func (au *Authenticator) GenerateToken(username string) (string, error) {
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(au.exp)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(au.secretKey)
	if err != nil {
		return "", entities.NewError(err, entities.Internal)
	}

	return tokenString, nil
}

func (au *Authenticator) Verify(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, entities.NewError(ErrMisMatchAlg, entities.UnAuthorize)
		}

		return au.secretKey, nil
	})
	if err != nil {
		return entities.NewError(err, entities.UnAuthorize)
	}

	if _, ok := token.Claims.(*Claims); ok && token.Valid {
		return nil
	}

	return entities.NewError(nil, entities.UnAuthorize)
}
