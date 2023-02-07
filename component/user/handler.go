package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/mehdieidi/storm/internal/derror"
	"github.com/mehdieidi/storm/model"
	"github.com/mehdieidi/storm/pkg/type/email"
	"github.com/mehdieidi/storm/pkg/type/file"
	"github.com/mehdieidi/storm/pkg/type/mobile"
	"github.com/mehdieidi/storm/pkg/type/password"
)

type createUserRequest struct {
	FirstName    string              `json:"first_name"`
	LastName     string              `json:"last_name"`
	Email        email.Email         `json:"email"`
	MobileNumber mobile.MobileNumber `json:"mobile_number"`
	Password     password.Password   `json:"password"`
	Avatar       *file.FileID        `json:"avatar"`
}

func (c createUserRequest) validate() error {
	if c.FirstName == "" {
		return derror.ErrInvalidFirstName
	}

	if c.LastName == "" {
		return derror.ErrInvalidLastName
	}

	if err := c.Email.Validate(); err != nil {
		return email.ErrInvalidEmail
	}

	if err := c.Password.Validate(); err != nil {
		return password.ErrInvalidPassword
	}

	return nil
}

func (h *httpHandler) createUserHandler(c echo.Context) error {
	var req createUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := req.validate(); err != nil {
		return err
	}

	u := model.User{
		Credential: model.Credential{
			FirstName:    req.FirstName,
			LastName:     req.LastName,
			Email:        req.Email,
			MobileNumber: req.MobileNumber,
			Password:     req.Password,
		},
		Avatar: req.Avatar,
	}

	u, err := h.userService.Create(c.Request().Context(), u)
	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, u)

	return nil
}
