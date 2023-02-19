package handler

import (
	"github.com/tahmooress/book-repository/entities"
	"github.com/tahmooress/book-repository/pkg/validator"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginRequest) validate() *Err {
	if l.Password == "" {
		return newErr(emptyPassword, nil)
	}

	_, err := validator.NewEmail(l.Email)
	if err != nil {
		return newErr(emailMalform, err)
	}

	return nil
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterRequest) validate() *Err {
	if r.Password == "" {
		return newErr(emptyPassword, nil)
	}

	if r.UserName == "" {
		return newErr(emptyUser, nil)
	}

	_, err := validator.NewEmail(r.Email)
	if err != nil {
		return newErr(emailMalform, err)
	}

	return nil
}

type RegisterResponse struct {
	ID    string `json:"id"`
	User  string `json:"user_name"`
	Email string `json:"email"`
}

type SearchBookResult struct {
	Books []*entities.Book `json:"books"`
}
