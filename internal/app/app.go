package app

import (
	"os"

	"github.com/V1merX/upserv-api/internal/config"
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
}
