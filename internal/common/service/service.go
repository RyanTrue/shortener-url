package service

import (
	"github.com/RyanTrue/shortener-url.git/internal/common/config"
	"github.com/RyanTrue/shortener-url.git/internal/common/repository"
	"github.com/jmoiron/sqlx"
)

type ServiceContainer struct {
	URL *urlService
	DB  *sqlx.DB
}

func NewServiceContainer(repo *repository.RepositoryContainer, config config.AppConfig) (*ServiceContainer, error) {
	URLService := urlService{
		Repo:   repo.URLrepo,
		config: config,
	}

	return &ServiceContainer{
		URL: &URLService,
		DB:  repo.Postgres,
	}, nil
}
