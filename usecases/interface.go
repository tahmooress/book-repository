package usecases

import (
	"context"

	"github.com/tahmooress/book-repository/entities"
)

type Service interface {
	entities.UserUseCase
	entities.BookUseCase
}

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	Set(ctx context.Context, user *entities.User) error
}

type UserAuthenticator interface {
	GenerateToken(username string) (string, error)
	Verify(tokenString string) error
}

type BookSearchService interface {
	SearchBook(ctx context.Context, query string) ([]*entities.Book, error)
}

type BookCacheService interface {
	Set(ctx context.Context, query string, book []*entities.Book) error
	GetByName(ctx context.Context, name string) ([]*entities.Book, error)
}
