package entities

import "context"

type Book struct {
	ID         string
	ImageName  string
	Title      string
	Content    string
	Slug       string
	Publishers []string
	Authors    []string
}

type BookUseCase interface {
	SearchBook(ctx context.Context, query string) ([]*Book, error)
}
