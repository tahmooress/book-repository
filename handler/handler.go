package handler

import (
	"net/http"

	"github.com/tahmooress/book-repository/logger"
	"github.com/tahmooress/book-repository/usecases"
)

type handler struct {
	srv    usecases.Service
	au     usecases.UserAuthenticator
	logger logger.Logger
}

type HandlerUseCase interface {
	Login() http.HandlerFunc
	Register() http.HandlerFunc
	SearchBook() http.HandlerFunc
	JSON(next http.Handler) http.Handler
	AuthenticateUser(next http.Handler) http.Handler
}

func New(srv usecases.Service, authService usecases.UserAuthenticator, logger logger.Logger) *handler {
	return &handler{
		srv:    srv,
		au:     authService,
		logger: logger,
	}
}
