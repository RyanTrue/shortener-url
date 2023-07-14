package main

import (
	"go.uber.org/zap"

	"github.com/RyanTrue/shortener-url.git/internal/common/config"
	"github.com/RyanTrue/shortener-url.git/internal/common/handler"
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

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	repo := make(map[string]string)

	storage, err := service.NewStorage(appConfig.Server.TempDirectory)
	if err != nil {
		panic(err)
	}

	err = storage.Read(&repo)
	if err != nil {
		panic(err)
	}

	db, err := repository.NewPostgresDB(appConfig.DataBase.ConnectionStr)
	if err != nil {
		panic(err)
	}

	PostgresRepo := repository.NewRepository(db)

	services, err := service.NewServiceContainer(repo, appConfig, storage, PostgresRepo)
	if err != nil {
		panic(err)
	}
	handler := handler.NewHandler(services)
	server := new(server.Server)

	if err := server.Run(appConfig.Server.ServerAddr, handler.InitRoutes(sugar)); err != nil {
		panic(err)
	}
}
