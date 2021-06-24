package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/mqtt-broker/package/service"
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
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.DELETE("/sign-out", h.signOut)
	}

	api := router.Group("/api", h.userIdentity)
	{
		profile := api.Group("/profile")
		{
			profile.GET("/me", h.getProfile)
		}
		mosquitto := api.Group("/mosquitto")
		{
			mosquitto.POST("/launch", h.MosquittoLaunch)
		}
		topics := api.Group("/topics")
		{
			topics.GET("/", h.getAllTopics)
			topics.DELETE("/:id", h.deleteTopic)
			topics.POST("/create", h.createTopic)
		}
	}
	return router
}
