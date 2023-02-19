package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tahmooress/book-repository/configs"
	"github.com/tahmooress/book-repository/handler"
	"github.com/tahmooress/book-repository/infrastructure/auth"
	"github.com/tahmooress/book-repository/infrastructure/cache"
	"github.com/tahmooress/book-repository/infrastructure/database"
	"github.com/tahmooress/book-repository/infrastructure/fidibo"
	"github.com/tahmooress/book-repository/infrastructure/port"
	"github.com/tahmooress/book-repository/logger"
	"github.com/tahmooress/book-repository/pkg/wrapper"
	"github.com/tahmooress/book-repository/usecases"
)

func main() {
	closer, errChan, err := runner()
	if err != nil {
		panic(err)
	}

	os.Exit(shutdown(errChan, closer))
}

// ruuner will instantiate and runn the program
// because main file dosent care about deffer when
// log.Fatal is calling which cause to resource leak.
func runner() (io.Closer, <-chan error, error) {
	ctx := context.Background()

	cfg, err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := logger.New(logger.Config{
		LogFilePath: cfg.LogPath,
		LogLevel:    cfg.LogLevel,
	})
	if err != nil {
		log.Fatal(err)
	}

	c := new(wrapper.Closer)

	c.Add(logger)

	defer func() {
		if err != nil {
			_ = c.Close()
		}
	}()

	db, err := database.New(ctx, database.Config{
		DBName:   cfg.DatabaseName,
		DBHost:   cfg.DatabaseIP,
		DBPort:   cfg.DatabasePort,
		User:     cfg.DatabaseUser,
		Password: cfg.DatabasePass,
	})
	if err != nil {
		log.Fatal(err)
	}

	c.Add(db)

	redis, err := cache.New(ctx, cache.Config{
		Host:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		TTL:      cfg.RedisTTL,
	})
	if err != nil {
		log.Fatal(err)
	}

	c.Add(redis)

	f := fidibo.New()

	service := usecases.New(db, f, redis)

	au := auth.New([]byte(cfg.JWTSecretKey), cfg.JWTExpiration)

	h := handler.New(service, au, logger)

	return port.NewHTTPServer(h, cfg)
}

func shutdown(errChan <-chan error, closer io.Closer) int {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	defer close(c)

	var existStatus int

	select {
	case <-c:
		err := closer.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "terminate signal received -> closer error: %s\n", err)

			existStatus = 1
		} else {
			fmt.Fprintf(os.Stdout, "terminate signal received -> shutdowned cleanly")
		}
	case err := <-errChan:
		fmt.Fprintln(os.Stderr, err)
	}

	return existStatus
}
