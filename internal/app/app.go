package app

import (
	"fmt"
	"log"

	"github.com/V1merX/upserv-api/internal/config"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(cfg.String())
}
