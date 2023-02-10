package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mehdieidi/storm/internal/domain"
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

func (u *user) createHandler(c echo.Context) error {
	// TODO
	c.String(200, "created")
	return nil
}

func (u *user) getHandler(c echo.Context) error {
	// TODO
	c.String(200, "get")

	return nil
}

func (u *user) listHandler(c echo.Context) error {
	// TODO
	c.String(200, "list")

	return nil
}

func (u *user) updateHandler(c echo.Context) error {
	// TODO
	c.String(200, "update")

	return nil
}

func (u *user) deleteHandler(c echo.Context) error {
	// TODO
	c.String(200, "delete")

	return nil
}
