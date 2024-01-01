package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"parking-lot/config"
	"parking-lot/internal/controller/rest"
	"parking-lot/pkg/httpserver"
	"parking-lot/pkg/logger"
	"parking-lot/pkg/mysql"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func Run() error {
	// initiate config
	config, err := config.NewConfig()
	if err != nil {
		return err
	}

	// initiate logger
	log, err := logger.New()
	if err != nil {
		return err
	}

	// initiate db
	db, err := mysql.ConnectDB(config.DB.URL)
	if err != nil {
		return err
	}
	defer db.Close()

	m, err := Migrate(context.Background(), config.DB.URL, log)
	if err != nil {
		return err
	}

	defer m.Close()

	// initiate persistence
	repoLayer := NewRepoLayer(db, log)

	// initiate service

	useCaseLayer := NewUsecaseLayer(repoLayer, log)

	// HTTP Server
	r := mux.NewRouter()
	rest.NewRouter(r, log, useCaseLayer.UserUseCase)
	httpServer := httpserver.New(r, httpserver.Port(config.Server.Port))

	log.Info(context.Background(), fmt.Sprintf("server started on port: %s", config.Server.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info(context.Background(), "app - Run - signal", zap.Any("signal", s.String()))
	case err = <-httpServer.Notify():
		log.Error(context.Background(), "app - Run - httpServer.Notify", zap.Error(err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(context.Background(), "app shutdown", zap.Error(err))
	}

	log.Info(context.Background(), "config initiated", zap.Any("config", config))
	return nil
}
