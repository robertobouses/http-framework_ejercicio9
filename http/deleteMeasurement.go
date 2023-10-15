package http

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) DeleteMeasurement(ctx *gin.Context, id string) {
	err := h.service.DeleteMeasurement(id)
	if err != nil {
		log.Print("Error al llamar de http a app. Función delete", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})

	}
	ctx.JSON(nethttp.StatusOK, "se ha borrado correctamente")
}
