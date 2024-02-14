package main

import (
	"log"

	scraper_app "github.com/fentezi/scraper"
	"github.com/fentezi/scraper/internal/handler"
)

func main() {
	srv := new(scraper_app.Server)
	hand := new(handler.Handler)
	if err := srv.Run("8080", hand.InitRouters()); err != nil {
		log.Fatal(err)
	}
}
