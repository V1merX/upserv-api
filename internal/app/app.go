package app

import (
	"os"

	"github.com/V1merX/upserv-api/internal/config"
	"github.com/V1merX/upserv-api/pkg/a2s"
	log "github.com/sirupsen/logrus"
)

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

	err = a2s.Run("95.213.255.150", 27415)
	if err != nil {
		log.Error(err)
	}
}
