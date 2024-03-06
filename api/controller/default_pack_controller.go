package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pancudaniel7/pack-sizes-service/api/dto"
	"github.com/pancudaniel7/pack-sizes-service/internal/service"
	"net/http"
	"strconv"
)

type DefaultPackController struct {
	service service.PackService
}

func NewDefaultPackController(service service.PackService, router *gin.Engine) *DefaultPackController {
	controller := &DefaultPackController{
		service: service,
	}
	controller.registerRoutes(router, controller)
	return controller
}

func (c *DefaultPackController) registerRoutes(router *gin.Engine, controller *DefaultPackController) {
	router.POST("/set-pack-sizes", controller.SetPackSize)
	router.GET("/calculate-packs", controller.CalculatePacks)
}

func (c *DefaultPackController) SetPackSize(ctx *gin.Context) {
	var packDTO dto.PackSizesDTO
	if err := ctx.ShouldBindJSON(&packDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.SetPackSize(packDTO); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *DefaultPackController) CalculatePacks(ctx *gin.Context) {
	orderQtyStr := ctx.Query("orderQty")
	orderQty, err := strconv.Atoi(orderQtyStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order quantity"})
		return
	}

	packs, err := c.service.CalculatePacks(orderQty)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, packs)
}
