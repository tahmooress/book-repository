package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tahmooress/book-repository/entities"
)

func (d *Database) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := `select * from users where email= $1`

	var user entities.User

	err := d.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, entities.NewError(err, entities.Internal)
	}

	return &user, nil
}

func (d *Database) Set(ctx context.Context, user *entities.User) error {
	query := `insert into users(id, name, email, password) values($1,$2,$3,$4)`

	_, err := d.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return entities.NewError(err, entities.Internal)
	}

	return nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
