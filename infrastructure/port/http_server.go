package port

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/tahmooress/book-repository/configs"
	"github.com/tahmooress/book-repository/handler"
	"github.com/tahmooress/book-repository/pkg/util"
)

const (
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 45 * time.Second
)

func NewHTTPServer(handler handler.HandlerUseCase, cfg *configs.AppConfigs) (
	io.Closer, <-chan error, error,
) {
	router := mux.NewRouter()

	router.Use(
		mux.MiddlewareFunc(handler.JSON),
	)

	router.HandleFunc("/register", handler.Register()).Methods(http.MethodPost)
	router.HandleFunc("/login", handler.Login()).Methods(http.MethodPost)

	authRouter := router.Methods(http.MethodPost, http.MethodGet).Subrouter()

	authRouter.Use(
		mux.MiddlewareFunc(handler.AuthenticateUser),
	)

	authRouter.Path("/search/books").Queries("keyword", "{key}").
		HandlerFunc(handler.SearchBook()).Methods(http.MethodPost)

	fmt.Println("server running on:", fmt.Sprintf("%s:%s", cfg.HTTPIP, cfg.HTTPPort))

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.HTTPIP, cfg.HTTPPort),
		Handler:      router,
		ReadTimeout:  util.ParseDurationWithDefault(cfg.HTTPReadTimeout, defaultReadTimeout),
		WriteTimeout: util.ParseDurationWithDefault(cfg.HTTPWriteTimeout, defaultWriteTimeout),
	}

	errChan := make(chan error)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Printf("handler: ListenAndServe() error: %s\n", err)

			errChan <- err
		}
	}()

	return srv, errChan, nil
}
