package users

import (
	"context"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, param CreateUserParam) (*User, error)
	FindUserByUsername(ctx context.Context, username string) (*User, error)
}

type CreateUserParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
