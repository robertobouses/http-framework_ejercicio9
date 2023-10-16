package http

import (
	"fmt"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

func (h *Http) PostMeasurementValue(ctx *gin.Context) {
	var measurementValueRequest entity.MeasurementValueRequest
	if err := ctx.ShouldBindJSON(&measurementValueRequest); err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error al hacer ShouldBindJSON": err.Error()})
		return
	}

	cubic, err := h.service.CalcCubic(measurementValueRequest.ID)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, err)
	}
	amount, err := h.service.PrintAmount(measurementValueRequest.NAME)

	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, err)
	}

	if float32(cubic) > amount {
		ctx.JSON(200, fmt.Sprintf("Los metros cúbicos %v son mayores que el importe %v", cubic, amount))
	} else {
		ctx.JSON(200, fmt.Sprintf("Los metros cúbicos %v son menores o iguales que el importe %v", cubic, amount))
	}
}
