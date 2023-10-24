package main

import (
	"PTIT_TN/cmd/serverd/routers"
	"PTIT_TN/internal/services/db"
	"PTIT_TN/internal/services/db/redis"
	"PTIT_TN/pkg"
	"PTIT_TN/pkg/config"
	log "PTIT_TN/pkg/logger"
	"PTIT_TN/pkg/rabbitMQ"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	ctx := context.TODO()
	ctxRb := context.Background()
	v, f := viper.New(), pflag.NewFlagSet(string(pkg.APIAppName), pflag.ExitOnError)
	cfg := config.New(v, f)

	logger := log.NewLogger(cfg.Log)
	log.SetStandardLogger(logger)

	logger.Info("Initialize database...")
	database := db.New(logger, cfg.Database)
	defer database.Close()

	logger.Info("Initialize redis...")
	redisCache, err := redis.New(ctx,
		fmt.Sprintf("%v:%v",
			viper.GetString("redis.host"),
			viper.GetString("redis.port")),
		logger)
	if err != nil {
		return err
	}

	logger.Info("Initialize rabbitMQ connection...")
	mqService, err := rabbitMQ.NewConnection(cfg.RabbitMq, logger, ctxRb)
	defer mqService.GetConnection().Close()
	if err != nil {
		return err
	}

	appEnv := viper.GetString("APPENV")
	addr := fmt.Sprintf(":%s", cfg.App.HttpAddr)

	server := &http.Server{
		Addr: addr,
		Handler: routers.New(
			logger,
			database,
			redisCache,
			mqService,
			appEnv == "dev"),
	}
	serverErr := make(chan error, 1) // Create a channel to receive server error
	go func() {
		logger.Info("Running server at: ", map[string]interface{}{
			"address": addr,
		})
		serverErr <- server.ListenAndServe()
	}()

	// Create a channel to receive OS signal
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM) // Notify when receive Interrupt or SIGTERM signal (cause by Ctrl + C)

	// Block until receive OS signal
	select {
	case <-shutdown:
		logger.Info("Shutting down server...")

		// Create context with timer for graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		// Ask server to shutdown gracefully else force shutdown
		if err := server.Shutdown(ctx); err != nil {
			logger.Error("could not stop server gracefully")
			return server.Close()
		}
		logger.Info("Shutdown server done!")
	case err := <-serverErr:
		return errors.WithStack(err)
	}

	return nil
}
