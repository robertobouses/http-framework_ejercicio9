package http

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) GetValue(ctx *gin.Context, name string) {
	value, err := h.service.PrintValue(name)
	if err != nil {
		log.Printf("Error al obtener valor", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}
	log.Print("valor", value)
	ctx.JSON(nethttp.StatusOK, value)
}
