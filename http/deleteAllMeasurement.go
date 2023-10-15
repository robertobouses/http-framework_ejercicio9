package http

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) DeleteAllMeasurement(ctx *gin.Context) {
	err := h.service.DeleteAllMeasurement()
	if err != nil {
		log.Print("Error al llamar http a app. Función delete", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(nethttp.StatusOK, "se ha borrado correctamente")
}
