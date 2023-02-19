package usecases

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/tahmooress/book-repository/entities"
	"golang.org/x/crypto/scrypt"
)

func hashPassword(password string) (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", entities.NewError(err, entities.Internal)
	}

	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", entities.NewError(err, entities.Internal)
	}

	hashedPW := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))

	return hashedPW, nil
}

func comparePasswords(storedPassword string, suppliedPassword string) (bool, error) {
	pwsalt := strings.Split(storedPassword, ".")

	if len(pwsalt) < 2 {
		return false, entities.NewError(nil, entities.UnAuthorize)
	}

	salt, err := hex.DecodeString(pwsalt[1])
	if err != nil {
		return false, entities.NewError(err, entities.UnAuthorize)
	}

	shash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)
	if err != nil {
		return false, entities.NewError(err, entities.UnAuthorize)
	}

	return hex.EncodeToString(shash) == pwsalt[0], nil
}
