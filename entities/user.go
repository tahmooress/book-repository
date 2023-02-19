package entities

import (
	"context"
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt *time.Time
}

type UserUseCase interface {
	RegisterUser(ctx context.Context, user *User) (*User, error)
	AuthenticateUser(ctx context.Context, user *User) (*User, error)
}
