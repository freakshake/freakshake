package controller

import (
	"strconv"

	"github.com/freakshake/logger"
	"github.com/freakshake/type/email"
	"github.com/freakshake/type/file"
	"github.com/freakshake/type/id"
	"github.com/freakshake/type/mobile"
	"github.com/freakshake/type/offlim"
	"github.com/freakshake/type/optional"
	"github.com/freakshake/type/password"
	"github.com/labstack/echo/v4"

	"github.com/freakshake/freakshake/internal/derror"
	"github.com/freakshake/freakshake/internal/domain"
	"github.com/freakshake/freakshake/pkg/response"
)

type user struct {
	userService domain.UserService
	logger      logger.Logger
}

func UserRoutes(g *echo.Group, userService domain.UserService, logger logger.Logger) {
	u := user{userService: userService, logger: logger}

	g.POST("/users", u.createHandler)
	g.GET("/users/:id", u.getHandler)
	g.GET("/users", u.listHandler)
	g.PATCH("/users", u.updateHandler)
	g.DELETE("/users/:id", u.deleteHandler)
}

type createUserRequest struct {
	Avatar       optional.Optional[file.ID] `json:"avatar" swaggertype:"string"`
	FirstName    string                     `json:"first_name"`
	LastName     string                     `json:"last_name"`
	Email        email.Email                `json:"email"`
	MobileNumber mobile.Number              `json:"mobile_number"`
	Password     password.Password          `json:"password"`
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

//	@Summary		new user
//	@Description	create a user.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		createUserRequest	true	"User"
//	@Success		200		{object}	domain.User
//	@Failure		400		{object}	map[string]string{error=string}	"Invalid request"
//	@Failure		500		{object}	map[string]string{error=string}	"Internal server error"
//	@Router			/users [post]
func (u user) createHandler(c echo.Context) error {
	var req createUserRequest
	if err := c.Bind(&req); err != nil {
		u.logger.Error(domain.UserDomain, logger.TransportLayer, err, logger.Args{})
		return derror.EncodeHTTPError(c.Response().Writer, err)
	}

	if err := req.validate(); err != nil {
		u.logger.Error(domain.UserDomain, logger.TransportLayer, err, logger.Args{})
		return derror.EncodeHTTPError(c.Response().Writer, err)
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
		u.logger.Error(domain.UserDomain, logger.TransportLayer, err, logger.Args{})
		return derror.EncodeHTTPError(c.Response().Writer, err)
	}

	u.logger.Info(domain.UserDomain, logger.TransportLayer, logger.Args{"id": user.ID})

	return response.EncodeHTTP(c.Response(), user)
}

//	@Summary		get user by id
//	@Description	get user by id.
//	@Tags			User
//	@Produce		json
//	@Param			id	path		uint	true	"User ID"
//	@Success		200	{object}	domain.User
//	@Failure		400	{object}	map[string]string{error=string}	"Invalid request"
//	@Failure		500	{object}	map[string]string{error=string}	"Internal server error"
//	@Failure		404	{object}	map[string]string{error=string}	"unknown user"
//	@Router			/users/:id [get]
func (u user) getHandler(c echo.Context) error {
	uID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		u.logger.Error(domain.UserDomain, logger.TransportLayer, err, logger.Args{})
		return derror.EncodeHTTPError(c.Response().Writer, err)
	}

	user, err := u.userService.Get(c.Request().Context(), id.ID[domain.User](uID))
	if err != nil {
		u.logger.Error(domain.UserDomain, logger.TransportLayer, err, logger.Args{})
		return derror.EncodeHTTPError(c.Response().Writer, err)
	}

	u.logger.Info(domain.UserDomain, logger.TransportLayer, logger.Args{"id": uID})

	return response.EncodeHTTP(c.Response(), user)
}

type listUsersRequest struct {
	Offset offlim.Offset `query:"offset"`
	Limit  offlim.Limit  `query:"limit"`
}

//	@Summary		list users
//	@Description	list users.
//	@Tags			User
//	@Produce		json
//	@Param			offset	query		uint	false	"Offset"
//	@Param			limit	query		uint	false	"Limit"
//	@Success		200		{object}	[]domain.User
//	@Failure		400		{object}	map[string]string{error=string}	"Invalid request"
//	@Failure		500		{object}	map[string]string{error=string}	"Internal server error"
//	@Router			/users [get]
func (u user) listHandler(c echo.Context) error {
	var req listUsersRequest
	if err := c.Bind(&req); err != nil {
		u.logger.Error(domain.UserDomain, logger.TransportLayer, err, logger.Args{})
		return derror.EncodeHTTPError(c.Response().Writer, err)
	}

	users, err := u.userService.List(c.Request().Context(), req.Offset, req.Limit)
	if err != nil {
		u.logger.Error(domain.UserDomain, logger.TransportLayer, err, logger.Args{})
		return derror.EncodeHTTPError(c.Response().Writer, err)
	}

	u.logger.Info(domain.UserDomain, logger.TransportLayer, logger.Args{"offset": req.Offset, "limit": req.Limit})

	return response.EncodeHTTP(c.Response(), users)
}

type updateUserRequest struct {
	ID           id.ID[domain.User]         `json:"id" swaggertype:"integer"`
	Avatar       optional.Optional[file.ID] `json:"avatar" swaggertype:"string"`
	FirstName    string                     `json:"first_name"`
	LastName     string                     `json:"last_name"`
	Email        email.Email                `json:"email"`
	MobileNumber mobile.Number              `json:"mobile_number"`
}

//	@Summary		update user
//	@Description	update a user.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		updateUserRequest	true	"User"
//	@Success		200		{object}	bool
//	@Failure		400		{object}	map[string]string{error=string}	"Invalid request"
//	@Failure		500		{object}	map[string]string{error=string}	"Internal server error"
//	@Failure		404		{object}	map[string]string{error=string}	"unknown user"
//	@Router			/users [patch]
func (u user) updateHandler(c echo.Context) error {
	var req updateUserRequest
	if err := c.Bind(&req); err != nil {
		u.logger.Error(domain.UserDomain, logger.TransportLayer, err, logger.Args{})
		return derror.EncodeHTTPError(c.Response().Writer, err)
	}

	user := domain.User{
		Avatar:       req.Avatar,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		MobileNumber: req.MobileNumber,
	}

	if err := u.userService.Update(c.Request().Context(), req.ID, user); err != nil {
		u.logger.Error(domain.UserDomain, logger.TransportLayer, err, logger.Args{})
		return derror.EncodeHTTPError(c.Response().Writer, err)
	}

	u.logger.Info(domain.UserDomain, logger.TransportLayer, logger.Args{"id": req.ID})

	return response.EncodeHTTP(c.Response(), true)
}

//	@Summary		delete user
//	@Description	delete a user.
//	@Tags			User
//	@Produce		json
//	@Param			id	path		uint	true	"User ID"
//	@Success		200	{object}	bool
//	@Failure		400	{object}	map[string]string{error=string}	"Invalid request"
//	@Failure		500	{object}	map[string]string{error=string}	"Internal server error"
//	@Failure		404	{object}	map[string]string{error=string}	"unknown user"
//	@Router			/users/:id [delete]
func (u user) deleteHandler(c echo.Context) error {
	uID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		u.logger.Error(domain.UserDomain, logger.TransportLayer, err, logger.Args{})
		return derror.EncodeHTTPError(c.Response().Writer, err)
	}

	if err := u.userService.Delete(c.Request().Context(), id.ID[domain.User](uID)); err != nil {
		u.logger.Error(domain.UserDomain, logger.TransportLayer, err, logger.Args{})
		return derror.EncodeHTTPError(c.Response().Writer, err)
	}

	u.logger.Info(domain.UserDomain, logger.TransportLayer, logger.Args{"id": uID})

	return response.EncodeHTTP(c.Response(), true)
}
