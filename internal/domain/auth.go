package domain

import (
	"context"

	"github.com/mehdieidi/freakshake/pkg/type/email"
	"github.com/mehdieidi/freakshake/pkg/type/password"
)

type AuthToken string

type LoginDTO struct {
	Email    email.Email       `json:"email"`
	Password password.Password `json:"password"`
}

type RegisterDTO struct {
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Email     email.Email       `json:"email"`
	Password  password.Password `json:"password"`
}

type AuthService interface {
	Login(context.Context, LoginDTO) (AuthToken, error)
	Register(context.Context, RegisterDTO) (User, error)
}
