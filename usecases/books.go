package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/book-repository/entities"
)

func (s *service) SearchBook(ctx context.Context, query string) ([]*entities.Book, error) {
	books, err := s.bc.GetByName(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("bookUseCase: RegisterUser >> %w", err)
	}

	if books != nil {
		return books, nil
	}

	books, err = s.bs.SearchBook(ctx, query)
	if err != nil {
		return nil, entities.NewError(err, entities.Internal)
	}

	err = s.bc.Set(ctx, query, books)
	if err != nil {
		return nil, entities.NewError(err, entities.Internal)
	}

	return books, nil
}
