package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/book-repository/entities"
	"github.com/tahmooress/book-repository/pkg/ulid"
)

func (s *service) RegisterUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	result, err := s.ur.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, fmt.Errorf("userUsecas: RegisterUser >> %w", err)
	}

	if result != nil {
		return nil, fmt.Errorf("userUsecas: RegisterUser >> %w",
			entities.NewError(nil, entities.DuplicateUser))
	}

	p, err := hashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("userUsecas: RegisterUser >> %w", err)
	}

	user.ID = ulid.Generate()
	user.Password = p

	err = s.ur.Set(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("userUsecas: RegisterUser >> %w", err)
	}

	return user, nil
}

func (s *service) AuthenticateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	result, err := s.ur.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, fmt.Errorf("userUsecas: AuthenticateUser >> %w", err)
	}

	if result == nil {
		return nil, fmt.Errorf("userUsecas: RegisterUser >> %w",
			entities.NewError(nil, entities.NotFound))
	}

	ok, err := comparePasswords(result.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("userUsecas: AuthenticateUser >> %w", err)
	}

	if !ok {
		return nil, fmt.Errorf("userUsecas: RegisterUser >> %w",
			entities.NewError(nil, entities.UnAuthorize))
	}

	return result, nil
}
