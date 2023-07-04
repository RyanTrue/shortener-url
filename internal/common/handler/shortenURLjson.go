package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/RyanTrue/shortener-url.git/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ShortenURLjson(c *gin.Context) {
	var req models.ShortenRequest
	body := c.Request.Body

	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			fmt.Printf("Failed to close body: %v", err)
		}
	}(body)

	dec := json.NewDecoder(body)
	if err := dec.Decode(&req); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if req.URL == "" {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	shortenURL := h.services.URL.ShortenURL(req.URL)

	res := models.ShortenResponce{
		Result: shortenURL,
	}

	c.JSON(http.StatusCreated, res)
}
