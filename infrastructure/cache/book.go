package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/tahmooress/book-repository/entities"
)

const bookPrefix = "book:"

func keyWithPrefix(key string) string { return fmt.Sprintf("%s%s", bookPrefix, key) }

func (c *Cache) Set(ctx context.Context, query string, books []*entities.Book) error {
	b, err := json.Marshal(books)
	if err != nil {
		return entities.NewError(err, entities.Internal)
	}

	return c.client.Set(ctx, keyWithPrefix(query), b, c.ttl).Err()
}

func (c *Cache) GetByName(ctx context.Context, name string) ([]*entities.Book, error) {
	result, err := c.client.Get(ctx, keyWithPrefix(name)).Bytes()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return nil, nil
		}

		return nil, entities.NewError(err, entities.Internal)
	}

	var books []*entities.Book

	err = json.Unmarshal(result, &books)
	if err != nil {
		return nil, entities.NewError(err, entities.Internal)
	}

	return books, nil
}

func (c *Cache) Close() error {
	return c.client.Close()
}
