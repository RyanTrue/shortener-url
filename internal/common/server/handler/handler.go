package handler

import (
	"github.com/RyanTrue/shortener-url.git/internal/common/storage"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *storage.ServiceContainer
}

func NewHandler(services *storage.ServiceContainer) *Handler {
	return &Handler{
		services: services,
	}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/", h.ShortenURL)
	router.GET("/:id", h.ExpandURL)

	return router
}
