package main

import (
	"flag"
	"github.com/RyanTrue/shortener-url.git/internal/common/config"
	"github.com/RyanTrue/shortener-url.git/internal/common/server"
	"github.com/RyanTrue/shortener-url.git/internal/common/server/handler"
	"github.com/RyanTrue/shortener-url.git/internal/common/storage"
	"log"
)

func main() {
	appConfig := config.AppConfig{}
	appConfig.InitAppConfig()
	flag.Parse()

	services := storage.NewServiceContainer(make(map[string]string), appConfig)
	handler := handler.NewHandler(services)
	server := &server.Server{}

	if err := server.Run(appConfig.Server.ServerAddr, handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
