package handler

import (
	"github.com/RyanTrue/shortener-url.git/internal/common/config"
	"github.com/RyanTrue/shortener-url.git/internal/common/middlewares"
	"github.com/RyanTrue/shortener-url.git/internal/common/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.ServiceContainer
	Cfg      config.AppConfig
}

func NewHandler(services *service.ServiceContainer, cfg config.AppConfig) *Handler {
	return &Handler{
		services: services,
		Cfg:      cfg,
	}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(middlewares.LogReqResInfo(), middlewares.DataCompressor())

	router.POST("/", h.ShortenURL)
	router.GET("/:id", h.ExpandURL)
	router.GET("/ping", h.pingDB)

	api := router.Group("/api")
	{
		api.POST("/shorten", h.ShortenURLjson)
		api.POST("/shorten/batch", h.batchURLinsert)
	}

	return router
}
