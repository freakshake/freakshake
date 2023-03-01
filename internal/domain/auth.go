package domain

import (
	"context"

	"github.com/freakshake/type/email"
	"github.com/freakshake/type/password"
)

const AuthDomain = "auth"

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
