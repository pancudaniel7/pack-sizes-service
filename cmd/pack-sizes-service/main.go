package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pancudaniel7/pack-sizes-service/api/controller"
	"github.com/pancudaniel7/pack-sizes-service/internal/dao"
	"github.com/pancudaniel7/pack-sizes-service/internal/model"
	"github.com/pancudaniel7/pack-sizes-service/internal/service"
	"github.com/pancudaniel7/pack-sizes-service/web/handlers"
	"github.com/spf13/viper"
)

func main() {
	initConfig()

	router := initRouter()
	initComponents(router)

	port := viper.GetString("app.port")
	host := viper.GetString("app.localhost")

	err := router.Run(host + ":" + port)
	if err != nil {
		panic(fmt.Errorf("Fatal error starting server: %w \n", err))
	}
}

func initComponents(router *gin.Engine) {
	inMemoryPackDao := dao.NewInMemoryPackDao(model.Pack{})
	packService := service.NewDefaultPackService(inMemoryPackDao)
	controller.NewDefaultPackController(packService, router)
	controller.NewHealthController(router)
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	})

	router.LoadHTMLGlob("web/templates/*.html")
	router.GET("/", handlers.IndexHandler)
	router.GET("/assets/css/styles.css", handlers.IndexCssHandler)
	router.GET("/assets/js/calculator.js", handlers.IndexJsHandler)

	return router
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
