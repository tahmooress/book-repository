package validator

import (
	"errors"
	"regexp"
	"strings"
)

var ErrInvalidEmail = errors.New("email format is invalid")

func NewEmail(v string) (string, error) {
	v = sanitizeEmail(v)

	return validateEmail(v)
}

func validateEmail(v string) (string, error) {
	if v == "" {
		return "", nil
	}

	re := regexp.MustCompile("^[a-z0-9.!#$%&'*+/=?^_`{|}~-]+" + "@" +
		"[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?(?:\\.[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)*$")
	if re.MatchString(v) {
		return v, nil
	}

	return "", ErrInvalidEmail
}

func sanitizeEmail(v string) string {
	v = strings.ToLower(v)
	v = strings.TrimSpace(v)

	return v
}
