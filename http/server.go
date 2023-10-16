package http

import (
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/http-framework_ejercicio9/app"
)

type HTTP interface {
	PostMeasurement(ctx *gin.Context)
	GetMeasurement(ctx *gin.Context)
	DeleteMeasurement(ctx *gin.Context, id string)
	GetCubic(ctx *gin.Context, id string)
	GetScale(ctx *gin.Context, id string)
	PostCompare(ctx *gin.Context)
	DeleteAllMeasurement(ctx *gin.Context)
	DeleteEmptyMeasurement(ctx *gin.Context)
	PostValue(ctx *gin.Context)
	GetValue(ctx *gin.Context, name string)
}

type Http struct {
	service app.APP
}

func NewHTTP(service app.APP) HTTP {
	return &Http{
		service: service,
	}
}
