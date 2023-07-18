package handler

import (
	"github.com/RyanTrue/shortener-url.git/internal/common/middlewares"
	"github.com/RyanTrue/shortener-url.git/internal/common/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.ServiceContainer
}

func NewHandler(services *service.ServiceContainer) *Handler {
	return &Handler{
		services: services,
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
	}

	return router
}
