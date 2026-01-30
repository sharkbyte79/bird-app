package service

import (
	"errors"
	"fmt"

	"github.com/sharkbyte79/birdup/internal/model"
	"github.com/sharkbyte79/birdup/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) (*UserService, error) {
	if r == nil {
		return nil, errors.New("failed to create User service")
	}
	return &UserService{repo: r}, nil
}

func (s *UserService) Create(u *model.User) error {
	if err := s.repo.Create(u); err != nil {
		return fmt.Errorf("error creating User record: %w", err)
	}
	return nil
}

// func (s *UserService) Read(firebaseID string) (*model.User, error) {}

// func (s *UserService) Update(firebaseID string) (*model.User, error) {}

// func (s *UserService) Delete(firebaseID string) (*model.User, error) {}
