package fidibo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/tahmooress/book-repository/entities"
)

const (
	address = "https://search.fidibo.com"
)

type FidiboSearchService struct {
	client http.Client
}

func New() *FidiboSearchService {
	return &FidiboSearchService{
		client: *http.DefaultClient,
	}
}

var ErrNotResponding = errors.New("provider not responding")

func (f *FidiboSearchService) SearchBook(ctx context.Context, query string) ([]*entities.Book, error) {
	data := url.Values{}
	data.Set("q", query)

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, address, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("FidiboSearchService: SearchBook() >> %w", err)
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := f.client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("FidiboSearchService: SearchBook() >> %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrNotResponding
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("FidiboSearchService: SearchBook() >> %w", err)
	}

	var books BooksResponse

	err = json.Unmarshal(b, &books)
	if err != nil {
		return nil, fmt.Errorf("FidiboSearchService: SearchBook() >> %w", err)
	}

	return adapter(books), nil
}

func adapter(input BooksResponse) []*entities.Book {
	books := make([]*entities.Book, len(input.Books.Hits.Hits))

	for i, obj := range input.Books.Hits.Hits {
		books[i] = &entities.Book{
			ID:         obj.Source.ID,
			ImageName:  obj.Source.ImageName,
			Title:      obj.Source.Title,
			Content:    obj.Source.Content,
			Slug:       obj.Source.Slug,
			Publishers: entities.Title(obj.Source.Publishers),
			Authors:    adaptAuthors(obj.Source.Authors),
		}
	}

	return books
}

func adaptAuthors(au []Author) []entities.Author {
	r := make([]entities.Author, len(au))

	for i := range au {
		r[i] = entities.Author{
			Name: au[i].Name,
		}
	}

	return r
}
