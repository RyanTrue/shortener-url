package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ExpandURL(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		http.Error(c.Writer, "Error reading id param", http.StatusInternalServerError)
		return
	}

	value, err := h.services.URL.ExpandURL(id)
	if err != nil {
		http.Error(c.Writer, "No original URL found", http.StatusNotFound)
		return
	}
	c.Status(http.StatusTemporaryRedirect)
	c.Header("Location", value)
}
