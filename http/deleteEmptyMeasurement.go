package http

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) DeleteEmptyMeasurement(ctx *gin.Context) {
	err := h.service.DeleteEmptyMeasurement()
	if err != nil {
		log.Print("Error al llamar http a app. Función delete", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(nethttp.StatusOK, "se han borrado los registros con partes vacías correctamente")
}
