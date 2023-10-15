package http

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) GetCubic(ctx *gin.Context, id string) {
	cubic, err := h.service.CalcCubic(id)
	if err != nil {
		log.Printf("Error", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(nethttp.StatusOK, cubic)

}
