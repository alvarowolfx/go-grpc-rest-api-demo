package users

import (
	"errors"

	"context"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(ctx context.Context, param CreateUserParam) (*User, error)
	CheckUser(ctx context.Context, username, password string) (*User, error)
}

type ServiceDeps struct {
	UserRepository Repository
}

type userService struct {
	ServiceDeps
}

func NewService(deps ServiceDeps) Service {
	return &userService{deps}
}

func (s *userService) CreateUser(ctx context.Context, param CreateUserParam) (*User, error) {
	user, err := s.UserRepository.FindUserByUsername(ctx, param.Username)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	encryptedPassword := string(hash)

	return s.UserRepository.CreateUser(ctx, CreateUserParam{
		Username: param.Username,
		Password: encryptedPassword,
	})
}

func (s *userService) CheckUser(ctx context.Context, username, password string) (*User, error) {
	user, err := s.UserRepository.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("username and password doesn't match")
	}

	return user, nil
}
