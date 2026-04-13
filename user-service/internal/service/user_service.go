package service

import (
	"errors"
	"user-service/internal/model"
	"user-service/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) Register(user model.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) Login(email, password string) (*model.User, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
