package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
	config Config
}

func New(config Config) *App {
	app := &App{
		rdb: redis.NewClient(&redis.Options{
			Addr: config.RedisAddress,
		}),
		config: config,
	}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to ping redis: %w", err)
	}
	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start the server: %w", err)
	}
	return nil
}
