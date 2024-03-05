package controller

import (
	"github.com/gin-gonic/gin"
)

type PackController interface {
	SetPackSize(ctx *gin.Context)
	CalculatePacks(ctx *gin.Context)
}
