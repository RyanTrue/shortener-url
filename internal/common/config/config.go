package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

type AppConfig struct {
	Server struct {
		DefaultAddr string `env:"BASE_URL"`
		ServerAddr  string `env:"SERVER_ADDRESS"`
		TempFolder  string `env:"FILE_STORAGE_PATH"`
	}
}

func (a *AppConfig) InitAppConfig() error {
	flag.StringVar(&a.Server.ServerAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&a.Server.DefaultAddr, "b", "http://localhost:8080", "default address and port of a shortened URL")
	flag.StringVar(&a.Server.TempFolder, "f", "/tmp/short-url-db.json", "default temp data storage path and filename")
	flag.Parse()

	err := env.Parse(a)
	if err != nil {
		return err
	}
	return nil
}
