package controller

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mehdieidi/storm/internal/derror"
	"github.com/mehdieidi/storm/internal/domain"
	"github.com/mehdieidi/storm/pkg/type/email"
	"github.com/mehdieidi/storm/pkg/type/file"
	"github.com/mehdieidi/storm/pkg/type/id"
	"github.com/mehdieidi/storm/pkg/type/mobile"
	"github.com/mehdieidi/storm/pkg/type/offlim"
	"github.com/mehdieidi/storm/pkg/type/password"
)

type user struct {
	UserService domain.UserService
}

func UserRoutes(e *echo.Echo, userService domain.UserService) {
	u := user{UserService: userService}

	e.POST("/users", u.createHandler)
	e.GET("/users/:id", u.getHandler)
	e.GET("/users", u.listHandler)
	e.PATCH("/users", u.updateHandler)
	e.DELETE("/users/:id", u.deleteHandler)
}

type createUserRequest struct {
	Avatar       *file.FileID        `json:"avatar"`
	FirstName    string              `json:"first_name"`
	LastName     string              `json:"last_name"`
	Email        email.Email         `json:"email"`
	MobileNumber mobile.MobileNumber `json:"mobile_number"`
	Password     password.Password   `json:"password"`
}

func (c createUserRequest) validate() error {
	if c.FirstName == "" {
		return derror.ErrInvalidFirstName
	}
	if c.LastName == "" {
		return derror.ErrInvalidLastName
	}
	if err := c.Email.Validate(); err != nil {
		return err
	}
	if err := c.MobileNumber.Validate(); err != nil {
		return err
	}
	if err := c.Password.Validate(); err != nil {
		return err
	}
	return nil
}

func (u *user) createHandler(c echo.Context) error {
	var req createUserRequest
	if err := c.Bind(&req); err != nil {
		return c.String(500, err.Error())
	}

	if err := req.validate(); err != nil {
		return c.String(500, err.Error())
	}

	user := domain.User{
		Avatar: req.Avatar,
		Credential: domain.Credential{
			FirstName:    req.FirstName,
			LastName:     req.LastName,
			Email:        req.Email,
			MobileNumber: req.MobileNumber,
			Password:     req.Password,
		},
	}

	user, err := u.UserService.Create(c.Request().Context(), user)
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(200, user)
}

func (u *user) getHandler(c echo.Context) error {
	uID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(500, err.Error())
	}

	user, err := u.UserService.Get(c.Request().Context(), id.ID[domain.User](uID))
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(200, user)
}

func (u *user) listHandler(c echo.Context) (err error) {
	o := c.QueryParam("offset")
	var offset int
	if o != "" {
		offset, err = strconv.Atoi(o)
		if err != nil {
			return c.String(500, err.Error())
		}
	}

	l := c.QueryParam("limit")
	limit := -1
	if l != "" {
		limit, err = strconv.Atoi(l)
		if err != nil {
			return c.String(500, err.Error())
		}
	}

	users, err := u.UserService.List(c.Request().Context(), offlim.Offset(offset), offlim.Limit(limit))
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(200, users)
}

func (u *user) updateHandler(c echo.Context) error {
	// TODO
	return nil
}

func (u *user) deleteHandler(c echo.Context) error {
	// TODO
	return nil
}
