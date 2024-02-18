package main

import (
	scraper_app "github.com/fentezi/scraper"
	"github.com/fentezi/scraper/internal/handler"
	"github.com/fentezi/scraper/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Start server")
	srv := new(scraper_app.Server)
	logger.Info("Init handler")
	hand := handler.NewHandler(logger)
	if err := srv.Run("8080", hand.InitRouters()); err != nil {
		logger.Fatal(err)
	}
}
