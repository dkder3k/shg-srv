package handler

import (
	"github.com/dkder3k/shg-srv/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		hosts := api.Group("/hosts")
		{
			hosts.GET("/", h.listHosts)
			hosts.POST("/", h.addHost)
			hosts.GET("/:id", h.getHost)
			hosts.DELETE("/:id", h.deleteHost)
			hosts.PUT("/:id", h.updateHost)

			hostConfigs := hosts.Group("/:id/configurations")
			{
				hostConfigs.GET("/", h.listHostConfigs)
				hostConfigs.POST("/", h.addHostConfig)
				hostConfigs.GET("/:confId", h.getHostConfig)
				hostConfigs.DELETE("/:confId", h.deleteHostConfig)
			}

			hostVars := hosts.Group("/:id/vars")
			{
				hostVars.GET("/", h.listHostVars)
				hostVars.POST("/", h.addHostVar)
				hostVars.GET("/:var", h.getHostVar)
				hostVars.DELETE("/:var", h.deleteHostVar)
				hostVars.PUT("/:var", h.updateHostVar)
			}
		}
		configurations := api.Group("/configurations")
		{
			configurations.GET("/", h.listConfigurations)
			configurations.POST("/", h.addConfiguration)
			configurations.GET("/:confId", h.getConfiguration)
			configurations.DELETE("/:confId", h.deleteConfiguration)
			configurations.PUT("/:confId", h.updateConfiguration)
		}
	}

	return router
}
