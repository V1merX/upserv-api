package app

import (
	"fmt"
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
		return
	}

	logLVL, err := log.ParseLevel(cfg.Logger.Level)
	if err != nil {
		panic(err)
	}

	log.SetLevel(logLVL)

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	fmt.Println(cfg.String())
}
