package handler

import (
	"flag"
	"github.com/RyanTrue/shortener-url.git/internal/common/config"
	"github.com/RyanTrue/shortener-url.git/internal/common/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestShortenURL(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	appConfig := config.AppConfig{}
	appConfig.InitAppConfig()
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
		test := test
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			с, _ := gin.CreateTestContext(w)

			с.Request, _ = http.NewRequest(test.method, test.url, strings.NewReader(test.body))

			h := Handler{
				services: storage.NewServiceContainer(testVault, appConfig),
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
