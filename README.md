# freakshake

freakshake is an opinionated guide and a demonstrative Go backend project. Attempting to structure the project in a way that complies with the SOLID principles, layered and clean architecture, and other software engineering best practices.

## Import Declarations

Order:

* stdlib
* 3rd party
* local

```go
import (
    "fmt"

    "github.com/labstack/echo/v4"

    "github.com/freakshake/freakshake/xsql"
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

Examples log output:

```log
{"level":"error","domain":"user","layer":"TRANSPORT","file":"/home/user/freakshake/internal/controller/user.go","line":76,"caller":"github.com/freakshake/internal/controller.user.createHandler","err":"code=400, message=Syntax error: offset=50, error=invalid character 's' after object key:value pair,"time":1677231468}
```

## Developing New Requirements

Now if you need to develop a new feature (after designing process):

* Create the domain model and interfaces in /internal/domain/
* Implement new domain's storage interface in the /internal/storage/ this should meet the storage needs of the domain
* Implement new domain's service interface in the /internal/service/ this should meet the business logic needs of the domain
* You should implement related defined errors along the way you develop the new feature in /internal/derror/ and add the new errors to the errStatusMap
* Implement handlers and http routes function in /internal/controller/ and call the function in the /cmd/service/main.go
* Add comments on top of the handlers for auto swagger doc generation
* Run $ make pipeline to start the service.
* Test endpoints
* Check swagger docs

## Errors

With no surprise errors are handled in the Go way!

```Go
if err != nil {
    return err
}
```

### Panics

In the main.go files we prefer to panic when we encounter an error. You can use the xerror package:

```go
xerror.PanicIf(err)
```

## When to consider abstraction?

In our opinion you should avoid abstraction-trap. Sure dependency inversion is cool, hiding detail is good, ... but we should consider that not always we need to abstract things using interface and stuff.

* For example when you are dealing with a cache, I mean the basic put, get, ... functionalities of a cache, there is a great chance where you are going to use this cache all over your project and if the cache tool you have used doesn't meet your requirements in the future, for let's say, performance reasons, you need to change code all over your project to replace the cache.
But if you have defined an interface and implemented that interface using a cache tool, the only place you needed to change was the implementation. Not all the places where the cache is used.

* But consider the case where you need to implement a token-based authentication feature having login and register endpoints. Now in the business logic layer you need to verify user credentials and create a token for the user. Thats where you probably are going to use something like JWT. So you need a JWT library to construct tokens. Now you DONT need to define abstractions and interfaces and make things more complex than it supposed to be. Because the only place you are using that lib is a well-known layer called business logic layer and its only ONE place. If you wanted to replace JWT with an other mechanism, say PASETO, you only need to change one implementation in only one layer!
