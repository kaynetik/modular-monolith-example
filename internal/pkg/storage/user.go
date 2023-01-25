package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/kaynetik/modular-monolith-example/internal/pkg/models"
	"github.com/uptrace/bun"
)

func (r *Repository) CreateUser(u *models.User) error {
	res, err := r.db.NewInsert().Model(u).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("an error occurred while creating a user: %w", err)
	}

	num, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("an error occurred while creating a user: %w", err)
	}

	if num == 0 {
		return errors.New("an error occurred while creating a user: no rows affected")
	}

	return nil
}

func (r *Repository) GetUserByID(id uuid.UUID) (*models.User, error) {
	user := new(models.User)

	err := r.db.NewSelect().Model(user).Where("? = ?", bun.Ident("id"), id).Limit(1).Scan(context.Background())
	if err != nil {
		return nil, fmt.Errorf("an error occurred while fetching user by ID: %w", err)
	}

	return user, nil
}

func (r *Repository) GetUserByEmail(e string) (*models.User, error) {
	return new(models.User), nil
}
