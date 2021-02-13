package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/go.git/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-up")
	}

	return router
}
