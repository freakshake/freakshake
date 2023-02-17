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
	"github.com/mehdieidi/storm/pkg/type/optional"
	"github.com/mehdieidi/storm/pkg/type/password"
)

type user struct {
	userService domain.UserService
}

func UserRoutes(g *echo.Group, userService domain.UserService) {
	u := user{userService: userService}

	g.POST("/users", u.createHandler)
	g.GET("/users/:id", u.getHandler)
	g.GET("/users", u.listHandler)
	g.PATCH("/users", u.updateHandler)
	g.DELETE("/users/:id", u.deleteHandler)
}

type createUserRequest struct {
	Avatar       optional.Optional[file.FileID] `json:"avatar"`
	FirstName    string                         `json:"first_name"`
	LastName     string                         `json:"last_name"`
	Email        email.Email                    `json:"email"`
	MobileNumber mobile.MobileNumber            `json:"mobile_number"`
	Password     password.Password              `json:"password"`
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
		Avatar:       req.Avatar,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		MobileNumber: req.MobileNumber,
		Password:     req.Password,
	}

	user, err := u.userService.Create(c.Request().Context(), user)
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

	user, err := u.userService.Get(c.Request().Context(), id.ID[domain.User](uID))
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(200, user)
}

type listUsersRequest struct {
	Offset offlim.Offset `query:"offset"`
	Limit  offlim.Limit  `query:"limit"`
}

func (u *user) listHandler(c echo.Context) (err error) {
	var req listUsersRequest
	if err := c.Bind(&req); err != nil {
		c.String(500, err.Error())
		return err
	}

	users, err := u.userService.List(c.Request().Context(), req.Offset, req.Limit)
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(200, users)
}

func (u *user) updateHandler(c echo.Context) error {
	// TODO
	panic("todo")
}

func (u *user) deleteHandler(c echo.Context) error {
	// TODO
	panic("todo")
}
