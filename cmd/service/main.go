package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mehdieidi/storm/config"
	"github.com/mehdieidi/storm/internal/controller"
	"github.com/mehdieidi/storm/internal/repository"
	"github.com/mehdieidi/storm/internal/usecase"
	"github.com/mehdieidi/storm/pkg/cache/redis"
	"github.com/mehdieidi/storm/pkg/mongo"
	"github.com/mehdieidi/storm/pkg/postgres"
	"github.com/mehdieidi/storm/pkg/xerror"
	httpserver "github.com/mehdieidi/storm/transport/server/http"
)

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

	cache := redis.New(cfg.Redis.Host, cfg.Redis.Port, redis.WithCredential(cfg.Redis.User, cfg.Redis.Password))

	mongoClient := mongo.NewClient(mongo.ConnectionString{
		Host:     cfg.Mongo.Host,
		Port:     cfg.Mongo.Port,
		User:     cfg.Mongo.User,
		Password: cfg.Mongo.Password,
		DB:       cfg.Mongo.DB,
	})
	xerror.PanicIf(mongo.Ping(mongoClient))
	defer mongoClient.Disconnect(context.TODO())

	// Storage (repository) layer.
	userPostgresStorage := repository.NewUserPostgresStorage(postgresDB)
	userMongoStorage := repository.NewUserMongoStorage(mongoClient)

	// Service (usecase) layer.
	userService := usecase.NewUserService(userPostgresStorage, userMongoStorage, cache)

	// HTTP Transport layer.
	e := echo.New()

	controller.UserRoutes(e, userService)

	errCh := make(chan error, 1)

	httpServer := httpserver.Listen(e, errCh, cfg.HTTPServer.IP, cfg.HTTPServer.Port)

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
		fmt.Println("[x] error", err)
	}
}
