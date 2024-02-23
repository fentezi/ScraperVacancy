package main

import (
	"flag"
	scraper_app "github.com/fentezi/scraper"
	"github.com/fentezi/scraper/internal/handler"
	"github.com/fentezi/scraper/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Start server")
	srv := new(scraper_app.Server)
	logger.Info("Init handler")
	port := flag.String("port", "8080", "port for server")
	flag.Parse()
	hand := handler.NewHandler(logger)
	if err := srv.Run(*port, hand.InitRouters()); err != nil {
		logger.Fatal(err)
	}
}
