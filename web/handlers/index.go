package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "index.html", nil)
}

func IndexCssHandler(c *gin.Context) {
	c.Header("Content-Type", "text/css")
	c.File("web/assets/css/styles.css")
}

func IndexJsHandler(c *gin.Context) {
	c.Header("Content-Type", "text/javascript")
	c.File("web/assets/js/calculator.js")
}
