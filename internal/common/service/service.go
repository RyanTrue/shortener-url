package service

import (
	"github.com/RyanTrue/shortener-url.git/internal/common/config"
)

type ServiceContainer struct {
	URL urlService
}

func NewServiceContainer(repo map[string]string, config config.AppConfig, storage *Storage) (*ServiceContainer, error) {
	URLService := urlService{
		repo:    repo,
		config:  config,
		storage: storage,
	}

	return &ServiceContainer{
		URL: URLService,
	}, nil
}
