package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/freakshake/cache/redis"
	"github.com/freakshake/logger/zerolog"
	"github.com/freakshake/mongo"
	"github.com/freakshake/postgres"
	"github.com/freakshake/xerror"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/freakshake/freakshake/api/swagger"
	"github.com/freakshake/freakshake/config"
	"github.com/freakshake/freakshake/internal/controller"
	"github.com/freakshake/freakshake/internal/service"
	"github.com/freakshake/freakshake/internal/storage"
	httpserver "github.com/freakshake/freakshake/transport/server/http"
)

//	@title		freakshake Backend Service
//	@BasePath	/api
func main() {
	cfg, err := config.Load()
	xerror.PanicIf(err)

	// Make peripheral connections. e.g. DB, cache, ...
	postgresDB, err := postgres.Open(postgres.DSN{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DB:       cfg.Postgres.DB,
	})
	xerror.PanicIf(err)
	xerror.PanicIf(postgres.Ping(postgresDB))
	defer postgresDB.Close()

	cache := redis.New(
		cfg.Redis.Host,
		cfg.Redis.Port,
		redis.WithCredential(cfg.Redis.User, cfg.Redis.Password),
	)
	xerror.PanicIf(cache.Ping())

	mongoClient, err := mongo.NewClient(mongo.ConnectionString{
		Host:     cfg.Mongo.Host,
		Port:     cfg.Mongo.Port,
		User:     cfg.Mongo.User,
		Password: cfg.Mongo.Password,
		DB:       cfg.Mongo.DB,
	})
	xerror.PanicIf(err)
	xerror.PanicIf(mongo.Ping(mongoClient))
	defer mongoClient.Disconnect(context.Background())

	logFile, err := os.OpenFile(cfg.Log.FileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	xerror.PanicIf(err)
	defer logFile.Close()

	logger := zerolog.New(logFile)

	// Auto migrate up.
	xerror.PanicIf(postgres.MigrateUp(postgresDB, cfg.Postgres.MigrationsPath))

	// Storage (repository) layer.
	userPostgresStorage := storage.NewUserPostgresStorage(postgresDB, logger)
	userMongoStorage := storage.NewUserMongoStorage(mongoClient, logger)

	// Service (usecase) layer.
	userService := service.NewUserService(userPostgresStorage, userMongoStorage, cache)
	authService := service.NewAuthService(cfg.Auth.SecretKey, cfg.Auth.TokenExpirationHours, userService)

	// HTTP transport layer.
	e := echo.New()

	g := e.Group("/api/v1")
	g.GET("/swagger/*", echoSwagger.WrapHandler)
	controller.UserRoutes(g, userService, logger)
	controller.AuthRoutes(g, authService, logger)

	httpServer := &http.Server{
		ReadTimeout:       cfg.HTTPServer.ReadTimeout * time.Second,
		WriteTimeout:      cfg.HTTPServer.WriteTimeout * time.Second,
		IdleTimeout:       cfg.HTTPServer.IdleTimeout * time.Second,
		ReadHeaderTimeout: cfg.HTTPServer.ReadHeaderTimeout * time.Second,
		Addr:              net.JoinHostPort(cfg.HTTPServer.IP, cfg.HTTPServer.Port),
		Handler:           e,
	}

	errCh := make(chan error, 1)

	httpserver.SpawnListener(httpServer, errCh)
	httpserver.SpawnShutdownListener(errCh, httpServer)

	if err := <-errCh; err != nil {
		fmt.Println("[x] error:", err)
	}
}
