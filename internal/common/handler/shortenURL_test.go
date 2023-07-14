package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/RyanTrue/shortener-url.git/internal/common/config"
	"github.com/RyanTrue/shortener-url.git/internal/common/service"
	"github.com/gin-gonic/gin"
)

func TestShortenURL(t *testing.T) {
	cfg := config.AppConfig{}
	cfg.Server.DefaultAddr = "http://localhost:8080"
	cfg.Server.ServerAddr = ":8080"
	cfg.Server.TempDirectory = "/tmp/short-url-db.json"

	storage, err := service.NewStorage(cfg.Server.TempDirectory)
	if err != nil {
		fmt.Printf("Failed to create test storage: %v", err)
		return
	}

	var testVault = make(map[string]string)
	type want struct {
		code     int
		response string
	}

	tests := []struct {
		name   string
		url    string
		method string
		body   string
		want   want
	}{
		{
			name:   "Test #1 - Regular URL",
			url:    "http://localhost:8080",
			method: "POST",
			body:   "https://yandex.ru",
			want: want{
				code:     201,
				response: "http://localhost:8080/e9db20b2",
			},
		},
		{
			name:   "Test #2 - Empty Body",
			url:    "http://localhost:8080",
			method: "POST",
			body:   "",
			want: want{
				code:     400,
				response: "",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			с, _ := gin.CreateTestContext(w)

			с.Request, _ = http.NewRequest(test.method, test.url, strings.NewReader(test.body))
			serviceContainer, err := service.NewServiceContainer(testVault, cfg, storage, nil)
			if err != nil {
				fmt.Printf("Error creating service container: %v", err)
			}
			h := Handler{
				services: serviceContainer,
			}
			h.ShortenURL(с)

			if с.Writer.Status() != test.want.code {
				t.Errorf("got status code %d, want %d", w.Code, test.want.code)
			}

			if body := strings.TrimSpace(w.Body.String()); body != test.want.response {
				t.Errorf("got response body '%s', want '%s'", body, test.want.response)
			}
		})
	}
}
