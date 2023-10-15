package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

func (h *Http) PostMeasurement(ctx *gin.Context) {
	var newMeasurement entity.Measurement
	err := ctx.BindJSON(&newMeasurement)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error al hacer BindJSON": err.Error()})
		return
	}

	err = h.service.CreateMeasurement(newMeasurement)
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	ctx.JSON(nethttp.StatusOK, newMeasurement)

}
