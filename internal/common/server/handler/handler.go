package handler

import (
	"github.com/RyanTrue/shortener-url.git/internal/common/storage"
	"github.com/gin-gonic/gin"
)

type logger interface {
	Infof(template string, args ...interface{})
}

type Handler struct {
	services *storage.ServiceContainer
}

func NewHandler(services *storage.ServiceContainer) *Handler {
	return &Handler{
		services: services,
	}
}

func (h Handler) InitRoutes(lg logger) *gin.Engine {
	router := gin.New()
	router.Use(h.logReqResInfo(lg))

	router.POST("/", h.ShortenURL)
	router.GET("/:id", h.ExpandURL)

	api := router.Group("/api")
	{
		api.POST("/shorten", h.ShortenURLjson)
	}

	return router
}
