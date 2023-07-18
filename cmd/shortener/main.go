package main

import (
	"github.com/RyanTrue/shortener-url.git/internal/common/config"
	"github.com/RyanTrue/shortener-url.git/internal/common/handler"
	"github.com/RyanTrue/shortener-url.git/internal/common/logger"
	"github.com/RyanTrue/shortener-url.git/internal/common/repository"
	"github.com/RyanTrue/shortener-url.git/internal/common/server"
	"github.com/RyanTrue/shortener-url.git/internal/common/service"

	_ "github.com/lib/pq"
)

func main() {
	appConfig := config.AppConfig{}
	err := appConfig.InitAppConfig()
	if err != nil {
		panic(err)
	}

	err = logger.InitLogger()
	if err != nil {
		panic(err)
	}

	repo, err := repository.NewRepositoryContainer(appConfig)
	if err != nil {
		panic(err)
	}

	services, err := service.NewServiceContainer(repo, appConfig)
	if err != nil {
		panic(err)
	}
	handler := handler.NewHandler(services)
	server := new(server.Server)

	if err := server.Run(appConfig.Server.ServerAddr, handler.InitRoutes()); err != nil {
		panic(err)
	}
}
