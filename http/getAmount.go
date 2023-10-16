package http

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) GetAmount(ctx *gin.Context, name string) {
	amount, err := h.service.PrintAmount(name)
	if err != nil {
		log.Printf("Error al obtener valor", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}
	log.Print("valor", amount)
	ctx.JSON(nethttp.StatusOK, amount)
}
