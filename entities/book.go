package entities

import "context"

type Book struct {
	ID         string
	ImageName  string
	Title      string
	Content    string
	Slug       string
	Publishers Title
	Authors    []Author
}

type Title struct {
	Title string
}

type Author struct {
	Name string
}

type BookUseCase interface {
	SearchBook(ctx context.Context, query string) ([]*Book, error)
}
