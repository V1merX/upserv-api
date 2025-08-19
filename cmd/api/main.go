package main

import (
	"log"

	"github.com/V1merX/upserv-api/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
