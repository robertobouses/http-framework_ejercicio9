package http

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

func (h *Http) PostCompare(ctx *gin.Context) {
	var measurementRequest entity.MeasurementRequest
	if err := ctx.ShouldBindJSON(&measurementRequest); err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error al hacer ShouldBindJSON": err.Error()})
		return
	}

	if len(measurementRequest.IDs) < 2 {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": "Se requieren al menos 2 IDs para comparar"})
		return
	}

	var cubicResults []int
	for _, id := range measurementRequest.IDs {
		cubic, err := h.service.CalcCubic(id)
		if err != nil {
			log.Printf("Error al calcular el cubo para ID %s: %v", id, err)
			ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		cubicResults = append(cubicResults, cubic)
	}

	minCubic, maxCubic := h.service.FindMinAndMaxCubic(cubicResults)

	ctx.JSON(nethttp.StatusOK, gin.H{
		"min_cubic": minCubic,
		"max_cubic": maxCubic,
	})
}
