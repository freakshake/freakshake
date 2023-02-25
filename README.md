# Storm

Storm is a template for a Go backend project. Attempting to structure the project in a way that complies with the SOLID principles and other software engineering best practices.

## Import Declarations

Import declarations order: 1. stdlib 2. 3rd party 3. local

```go
import (
    "fmt"

    "github.com/labstack/echo/v4"

    "github.com/freakshake/xsql
)
```

## Swagger Docs

We can use the <https://github.com/swaggo/swag> tool to automatically generate swagger specifications out the application comments.
These generated docs are located in the /api/swagger directory and the content of this directory are added to the .gitignore
file because they can get re-generated anytime.

### Can use the swagger target in the Makefile to generate docs

```shell
$ make swagger
...
```

### in /cmd/service/main.go

```go
// @title freakshake Backend Service
// @BasePath /api
func main() {
```

### in /internal/controller/user.go

Check the <https://github.com/swaggo/swag> docs for additional syntax. You can also check other handler comments located in /internal/controller/user.go

```go
// @Summary new user
// @Description create a new user.
// @Tags User
// @Accept json
// @Produce json
// @Param user body createUserRequest true "User"
// @Success 200 {object} domain.User
// @Failure 400 {object} map[string]string{error=string} "Invalid request"
// @Failure 500 {object} map[string]string{error=string} "Internal server error"
// @Router /users [post]
func (u user) createHandler(c echo.Context) (err error) {
```

* Formatting the swaggo comments and Go code are possible using the fmt Makefile target.

```shell
$ make fmt
...
```

## Config

The /config package is used for configuration management. Refer to the README.md file in /config for more detail.

## Errors

### Panics

In the main.go files we prefer to panic when we encounter an error. You can use the xerror package:

```go
xerror.PanicIf(err)
```
