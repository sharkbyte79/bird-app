package repository

import (
	"errors"
	"fmt"

	db "github.com/sharkbyte79/birdup/internal/database"
	models "github.com/sharkbyte79/birdup/internal/model"
)

type UserRepository interface {
	// Create inserts a new User record or returns an error.
	Create(u *models.User) error
	// GetOne returns a pointer to a User with a matching PK or an error.
	// GetOne(firebaseId string) (*models.User, error)
}

type Repository struct {
	Store *db.Store // Underlying database connection used by this repository
}

func NewRepository(store *db.Store) (UserRepository, error) {
	if store == nil {
		return nil, errors.New("failed to create new UserRepository")
	}
	return &Repository{Store: store}, nil
}

// Create inserts a new User record matching the given argument.
func (r *Repository) Create(u *models.User) error {
	_, err := r.Store.Db.Exec(
		"INSERT INTO users (firebaseId, email, createdAt) VALUES ($1, $2, $3);",
		u.FirebaseId,
		u.Email,
		u.CreatedAt)
	if err != nil {
		return fmt.Errorf("Error creating User record: %w", err)
	}
	return nil
}
