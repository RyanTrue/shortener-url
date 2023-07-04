package main

import (
	"flag"

	"github.com/RyanTrue/shortener-url.git/internal/common/config"
	"github.com/RyanTrue/shortener-url.git/internal/common/server"
	"github.com/RyanTrue/shortener-url.git/internal/common/server/handler"
	"github.com/RyanTrue/shortener-url.git/internal/common/storage"
	"go.uber.org/zap"
)

func main() {
	appConfig := config.AppConfig{}
	appConfig.InitAppConfig()
	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	repo := make(map[string]string)

	services := storage.NewServiceContainer(repo, appConfig)
	handler := handler.NewHandler(services)
	server := new(server.Server)

	if err := server.Run(appConfig.Server.ServerAddr, handler.InitRoutes(sugar)); err != nil {
		panic(err)
	}
}
