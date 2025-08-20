package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/V1merX/upserv-api/internal/config"
	"github.com/V1merX/upserv-api/internal/server"
	"github.com/V1merX/upserv-api/internal/service"
	transport "github.com/V1merX/upserv-api/internal/transport/rest"
	auth "github.com/V1merX/upserv-api/pkg/auth/jwt"
	log "github.com/sirupsen/logrus"
)

const timeout = 5 * time.Second

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
	}

	logLVL, err := log.ParseLevel(cfg.Logger.Level)
	if err != nil {
		panic(err)
	}

	log.SetLevel(logLVL)

	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		log.Error(err)
		return
	}

	services := service.NewServices()
	handlers := transport.NewHandler(services, *tokenManager)

	srv := server.NewServer(cfg, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	log.Info("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Errorf("failed to stop server: %v", err)
	}
}
