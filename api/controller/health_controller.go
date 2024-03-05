package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

func NewHealthController(router *gin.Engine) *HealthController {
	controller := &HealthController{}
	controller.registerRoutes(router, controller)
	return controller
}

func (c *HealthController) registerRoutes(router *gin.Engine, controller *HealthController) gin.IRoutes {
	return router.GET("/health", controller.CheckHealth)
}

func (c *HealthController) CheckHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "Up!"})
}
