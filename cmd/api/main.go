package main

import (
	"github.com/V1merX/upserv-api/internal/app"
)

const configsDir = "config"

func main() {
	app.Run(configsDir)
}
