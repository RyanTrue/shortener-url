package storage

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/RyanTrue/shortener-url.git/internal/common/config"
)

type urlService struct {
	repo   map[string]string
	config config.AppConfig
}

func (u *urlService) ShortenURL(body string) string {
	hasher := md5.New()
	hasher.Write([]byte(body))
	hash := hex.EncodeToString(hasher.Sum(nil))[:8]
	shortURL := fmt.Sprintf("%s/%s", u.config.Server.DefaultAddr, hash)
	if _, ok := u.repo[hash]; !ok {
		u.repo[hash] = body
	}
	return shortURL
}

func (u *urlService) ExpandURL(path string) (string, error) {
	if value, ok := u.repo[path]; ok {
		return value, nil
	} else {
		return "", fmt.Errorf("URL path '%s' not found", path)
	}
}
