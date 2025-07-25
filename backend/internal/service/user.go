package service

import (
	"errors"
	"glasscutting/internal/domain/model"
	"glasscutting/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Register(name, email, password, role string) (*model.User, error) {
	u := &model.User{Name: name, Email: email, Password: password, Role: role}
	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserService) Login(email, password string) (*model.User, error) {
	u, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if u.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return u, nil
}
