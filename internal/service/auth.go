package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/mehdieidi/freakshake/internal/derror"
	"github.com/mehdieidi/freakshake/internal/domain"
)

type auth struct {
	secretKey            string
	tokenExpirationHours uint
	userService          domain.UserService
}

func NewAuthService(
	secretKey string,
	tokenExpirationHours uint,
	userService domain.UserService,
) domain.AuthService {
	return &auth{
		secretKey:            secretKey,
		tokenExpirationHours: tokenExpirationHours,
		userService:          userService,
	}
}

func (s *auth) Login(ctx context.Context, l domain.LoginDTO) (domain.AuthToken, error) {
	u, err := s.userService.GetByEmail(ctx, l.Email)
	if err != nil {
		return "", derror.ErrInvalidCredentials
	}

	if !l.Password.CompareWithHashedPassword(u.HashedPassword) {
		return "", derror.ErrInvalidCredentials
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":  u.ID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(s.tokenExpirationHours) * time.Hour).Unix(),
	})

	token, err := t.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return domain.AuthToken(token), nil
}

func (s *auth) Register(ctx context.Context, r domain.RegisterDTO) (domain.User, error) {
	if _, err := s.userService.GetByEmail(ctx, r.Email); err != nil {
		if !errors.Is(err, derror.ErrUnknownUser) {
			return domain.User{}, err
		}
	} else {
		return domain.User{}, derror.ErrUserExists
	}

	user := domain.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Password:  r.Password,
	}

	user, err := s.userService.Create(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
