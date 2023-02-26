package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/freakshake/cache/redis"
	"github.com/freakshake/logger/zerolog"
	"github.com/freakshake/postgres"
	"github.com/freakshake/xerror"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/freakshake/api/swagger"
	"github.com/freakshake/config"
	"github.com/freakshake/internal/controller"
	"github.com/freakshake/internal/service"
	"github.com/freakshake/internal/storage"
	"github.com/freakshake/pkg/mongo"
	httpserver "github.com/freakshake/transport/server/http"
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
	defer postgresDB.Close()

	cache := redis.New(
		cfg.Redis.Host,
		cfg.Redis.Port,
		redis.WithCredential(cfg.Redis.User, cfg.Redis.Password),
	)

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

	// Auto migrate up.
	xerror.PanicIf(postgres.MigrateUp(postgresDB, cfg.Postgres.MigrationsPath))

	logger := zerolog.New(logFile)

	// Storage (repository) layer.
	userPostgresStorage := storage.NewUserPostgresStorage(postgresDB, logger)
	userMongoStorage := storage.NewUserMongoStorage(mongoClient)

	// Service (usecase) layer.
	userService := service.NewUserService(userPostgresStorage, userMongoStorage, cache)
	authService := service.NewAuthService(cfg.Auth.SecretKey, cfg.Auth.TokenExpirationHours, userService)

	// HTTP transport layer.
	e := echo.New()

	g := e.Group("/api/v1")
	g.GET("/swagger/*", echoSwagger.WrapHandler)
	controller.UserRoutes(g, userService, logger)
	controller.AuthRoutes(g, authService)

	errCh := make(chan error, 1)

	httpServer := &http.Server{
		ReadTimeout:       cfg.HTTPServer.ReadTimeout * time.Second,
		WriteTimeout:      cfg.HTTPServer.WriteTimeout * time.Second,
		IdleTimeout:       cfg.HTTPServer.IdleTimeout * time.Second,
		ReadHeaderTimeout: cfg.HTTPServer.ReadHeaderTimeout * time.Second,
		Addr:              net.JoinHostPort(cfg.HTTPServer.IP, cfg.HTTPServer.Port),
		Handler:           e,
	}

	httpserver.SpawnListener(httpServer, errCh)

	// Graceful shutdown in case of common signals.
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-ctx.Done()

		fmt.Println("\n[-] Shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer func() {
			stop()
			cancel()
			close(errCh)
		}()

		httpServer.SetKeepAlivesEnabled(false)

		if err := httpServer.Shutdown(ctxTimeout); err != nil {
			errCh <- err
		}

		fmt.Println("[-] Shutdown completed")
	}()

	if err := <-errCh; err != nil {
		fmt.Println("[x] error:", err)
	}
}
