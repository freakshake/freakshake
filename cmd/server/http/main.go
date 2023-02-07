package main

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/mehdieidi/storm/infrastructure/mongo"
	"github.com/mehdieidi/storm/infrastructure/postgres"
	"github.com/mehdieidi/storm/pkg/xerror"

	"github.com/mehdieidi/storm/component/user"
)

func main() {
	postgresDB, err := postgres.Open(postgres.DSN{
		Host:     "todo",
		Port:     "todo",
		User:     "todo",
		DB:       "todo",
		Password: "todo",
	})
	xerror.PanicIf(err)
	defer postgresDB.Close()

	mongoClient := mongo.NewClient(mongo.ConnectionString{
		Host: "todo",
		Port: "todo",
	})
	defer mongoClient.Disconnect(context.TODO())

	xerror.PanicIf(mongo.Ping(mongoClient))

	// Storage layer.
	userPostgresStorage := user.NewPostgresStorage(postgresDB)
	userMongoStorage := user.NewMongoStorage(mongoClient)

	// Service layer.
	userService := user.NewService(userPostgresStorage, userMongoStorage)

	// HTTP layer.
	e := echo.New()
	g := e.Group("/api/v1/")

	user.Routes(g, userService)

	e.Logger.Fatal(e.Start("todo"))
}
