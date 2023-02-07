package user

import (
	"github.com/labstack/echo/v4"
	"github.com/mehdieidi/storm/model"
)

type httpHandler struct {
	userService model.UserService
}

func Routes(e *echo.Group, userService model.UserService) {
	h := httpHandler{userService: userService}

	e.POST("/", h.createUserHandler)
}
