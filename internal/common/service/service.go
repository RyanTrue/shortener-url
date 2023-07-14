package service

import (
	"github.com/RyanTrue/shortener-url.git/internal/common/config"
	"github.com/RyanTrue/shortener-url.git/internal/common/repository"
)

type ServiceContainer struct {
	URL urlService
	DB  *repository.Repository
}

func NewServiceContainer(repo map[string]string, config config.AppConfig, storage *Storage, db *repository.Repository) (*ServiceContainer, error) {
	URLService := urlService{
		repo:    repo,
		db:      db,
		config:  config,
		storage: storage,
	}

	return &ServiceContainer{
		URL: URLService,
		DB:  db,
	}, nil
}
