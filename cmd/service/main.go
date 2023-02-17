package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/mehdieidi/storm/api/swagger"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/mehdieidi/storm/config"
	"github.com/mehdieidi/storm/internal/controller"
	"github.com/mehdieidi/storm/internal/service"
	"github.com/mehdieidi/storm/internal/storage"
	"github.com/mehdieidi/storm/pkg/cache/redis"
	"github.com/mehdieidi/storm/pkg/mongo"
	"github.com/mehdieidi/storm/pkg/postgres"
	"github.com/mehdieidi/storm/pkg/xerror"
	httpserver "github.com/mehdieidi/storm/transport/server/http"
)

// @title    Medad Backend Service
// @BasePath /api
func main() {
	cfg, err := config.Read()
	xerror.PanicIf(err)

	postgresDB, err := postgres.Open(postgres.DSN{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DB:       cfg.Postgres.DB,
	})
	xerror.PanicIf(err)
	defer postgresDB.Close()

	xerror.PanicIf(postgres.MigrateUp(postgresDB, cfg.Postgres.MigrationsPath))

	cache := redis.New(cfg.Redis.Host, cfg.Redis.Port, redis.WithCredential(cfg.Redis.User, cfg.Redis.Password))

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

	// Storage (repository) layer.
	userPostgresStorage := storage.NewUserPostgresStorage(postgresDB)
	userMongoStorage := storage.NewUserMongoStorage(mongoClient)

	// Service (usecase) layer.
	userService := service.NewUserService(userPostgresStorage, userMongoStorage, cache)

	// HTTP transport layer.
	e := echo.New()

	g := e.Group("/api/v1")
	g.GET("/swagger/*", echoSwagger.WrapHandler)
	controller.UserRoutes(g, userService)

	errCh := make(chan error, 1)

	httpServer := httpserver.Listen(e, errCh, cfg.HTTPServer.IP, cfg.HTTPServer.Port)

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
