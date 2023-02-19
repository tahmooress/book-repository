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

	var result BooksResponse

	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, fmt.Errorf("FidiboSearchService: SearchBook() >> %w", err)
	}

	fmt.Printf("%#+v ", result)

	return nil, nil
}
