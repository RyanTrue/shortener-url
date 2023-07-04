package handler

import (
	"flag"
	"fmt"
	"github.com/RyanTrue/shortener-url.git/internal/common/config"
	"github.com/RyanTrue/shortener-url.git/internal/common/storage"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestShortenURLjson(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	appConfig := config.AppConfig{}
	appConfig.InitAppConfig()
	testVault := make(map[string]string)
	type want struct {
		code     int
		response map[string]string
	}

	tests := []struct {
		name   string
		url    string
		method string
		body   map[string]string
		want   want
	}{
		{
			name:   "Test #1 - Regular URL",
			url:    "http://localhost:8080",
			method: "POST",
			body: map[string]string{
				"url": "https://yandex.ru",
			},
			want: want{
				code: 201,
				response: map[string]string{
					"result": "http://localhost:8080/e9db20b2",
				},
			},
		},
		{
			name:   "Test #2 - Empty Body",
			url:    "http://localhost:8080",
			method: "POST",
			body:   map[string]string{},
			want: want{
				code:     400,
				response: map[string]string{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			bodyStr, err := json.Marshal(test.body)
			if err != nil {
				fmt.Println("Error marshalling body:", err)
			}

			c.Request, _ = http.NewRequest(test.method, test.url, strings.NewReader(string(bodyStr)))
			h := Handler{
				services: storage.NewServiceContainer(testVault, appConfig),
			}
			h.ShortenURLjson(c)

			if c.Writer.Status() != test.want.code {
				t.Errorf("got status code %d, want %d", w.Code, test.want.code)
			}
		})
	}
}
