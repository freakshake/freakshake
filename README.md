# freakshake

freakshake is a guide and a demonstrative Go backend project. Attempting to structure the project in a way that complies with the SOLID principles and other software engineering best practices.

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

## Database

For each database e.g. postgres, mongo,... we should implement a constructor package like <https://github.com/freakshake/postgres>
or <https://github.com/freakshake/mongo> in the /pkg directory. We should implement a struct for a DSN or connection string or something like this.

We should ping the DB if possible for a quick health check.

## Logger

Currently logger component is using the <https://github.com/freakshake/logger> Logger interface and implements zerolog for structured logging. But the destination of logs are a single file on the same server as the backend app. This is bad :/

* Todo: Check proposal for implementing an event-based logging system using message queues, workers, remote server, and
log transport protocols.

Logging convention we use is logging all the errors in the storage layer and controller layer when binding request, validating, and calling the service layer.
Only log the ID of the entity being requested and processed as an information log in the corresponding controller. We do not log the entire object.

Examples of logs:

```log
{"level":"error","domain":"user","layer":"TRANSPORT","file":"/home/user/freakshake/internal/controller/user.go","line":76,"caller":"github.com/freakshake/internal/controller.user.createHandler","err":"code=400, message=Syntax error: offset=50, error=invalid character 's' after object key:value pair,"time":1677231468}
```

## Errors

### Panics

In the main.go files we prefer to panic when we encounter an error. You can use the xerror package:

```go
xerror.PanicIf(err)
```
