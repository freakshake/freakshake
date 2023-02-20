package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mehdieidi/storm/internal/domain"
)

type auth struct {
	authService domain.AuthService
}

func AuthRoutes(g *echo.Group, authService domain.AuthService) {
	a := auth{authService: authService}

	g.POST("/auth/login", a.LoginHandler)
	g.POST("/auth/register", a.RegisterHandler)
}

// @Summary     login user
// @Description login user.
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       login body     domain.LoginDTO true "login"
// @Success     200   {object} domain.AuthToken
// @Failure     400   {object} map[string]string{error=string} "Invalid request"
// @Failure     500   {object} map[string]string{error=string} "Internal server error"
// @Router      /auth/login [post]
func (a auth) LoginHandler(c echo.Context) error {
	var req domain.LoginDTO
	if err := c.Bind(&req); err != nil {
		c.String(500, err.Error())
		return err
	}

	token, err := a.authService.Login(c.Request().Context(), req)
	if err != nil {
		c.String(500, err.Error())
		return err
	}

	return c.JSON(200, token)
}

// @Summary     register a user
// @Description register a user.
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       register body     domain.RegisterDTO true "register"
// @Success     200      {object} domain.User
// @Failure     400      {object} map[string]string{error=string} "Invalid request"
// @Failure     500      {object} map[string]string{error=string} "Internal server error"
// @Router      /auth/register [post]
func (a auth) RegisterHandler(c echo.Context) error {
	var req domain.RegisterDTO
	if err := c.Bind(&req); err != nil {
		c.String(500, err.Error())
		return err
	}

	user, err := a.authService.Register(c.Request().Context(), req)
	if err != nil {
		c.String(500, err.Error())
		return err
	}

	return c.JSON(200, user)
}
