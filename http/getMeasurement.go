package http

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) GetMeasurement(ctx *gin.Context) {
	measurement, err := h.service.PrintMeasurement()
	if err != nil {
		log.Printf("Error al obtener medición", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.Print("medición en cada capa http:", measurement)
	ctx.JSON(nethttp.StatusOK, measurement)
}
