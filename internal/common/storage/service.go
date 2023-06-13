package storage

import (
	"github.com/RyanTrue/shortener-url.git/internal/common/config"
)

type ServiceContainer struct {
	URL urlService
}

func NewServiceContainer(repo map[string]string, config config.AppConfig) *ServiceContainer {
	return &ServiceContainer{
		URL: urlService{
			repo:   repo,
			config: config,
		},
	}
}
