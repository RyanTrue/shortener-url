package handlers

import (
	"fmt"
	"github.com/RyanTrue/shortener-url.git/internal/app/tools"
	"io"
	"net/http"
	"path"
)

type URLData struct {
	URL string
}

var vault = make(map[string]string)

func HTTPHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(res, "Error parsing params", http.StatusInternalServerError)
	}
	switch req.Method {
	case "POST":
		ShortenURL(res, req)
	case "GET":
		GetOriginalURL(res, req)
	}
}

func ShortenURL(res http.ResponseWriter, req *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing Body: %v", err)
		}
	}(req.Body)

	if req.Method != http.MethodPost && req.Method != http.MethodGet {
		http.Error(res, "Only POST and GET methods allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error reading request body", http.StatusInternalServerError)
		return
	}

	bodyStr := string(body)
	if bodyStr == "" {
		http.Error(res, "Can't shorten empty string", http.StatusBadRequest)
		return
	}

	shortURL, hash := tools.Shorten(bodyStr)
	if _, ok := vault[hash]; !ok {
		vault[hash] = bodyStr
	}
	res.WriteHeader(http.StatusCreated)
	if _, err = res.Write([]byte(shortURL)); err != nil {
		http.Error(res, "Error sending response", http.StatusInternalServerError)
	}
}

func GetOriginalURL(res http.ResponseWriter, req *http.Request) {
	_, id := path.Split(req.URL.Path)
	fmt.Println(vault)
	if value, ok := vault[id]; ok {
		res.Header().Set("Location", value)
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Error(res, "URL not found", http.StatusNotFound)
	}
}
