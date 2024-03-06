package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pancudaniel7/pack-sizes-service/api/controller"
	"github.com/pancudaniel7/pack-sizes-service/internal/dao"
	"github.com/pancudaniel7/pack-sizes-service/internal/model"
	"github.com/pancudaniel7/pack-sizes-service/internal/service"
)

func main() {
	router := initRouter()

	inMemoryPackDao := dao.NewInMemoryPackDao(model.Pack{})
	packService := service.NewDefaultPackService(inMemoryPackDao)
	controller.NewDefaultPackController(packService, router)
	controller.NewHealthController(router)

	router.Run()
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	})
	return router
}
