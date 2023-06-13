package handler

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func (h *Handler) ShortenURL(c *gin.Context) {
	body := c.Request.Body

	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Fatal(err)
		}
	}(body)

	data, err := io.ReadAll(body)
	if err != nil {
		http.Error(c.Writer, "Error reading request body", http.StatusInternalServerError)
		return
	}
	if len(data) == 0 {
		http.Error(c.Writer, "", http.StatusBadRequest)
		return
	}

	bodyStr := string(data)
	shortURL := h.services.URL.ShortenURL(bodyStr)

	c.String(http.StatusCreated, shortURL)
}
