package http

import (
	"fmt"
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) GetScale(ctx *gin.Context, id string) {
	cubic, err := h.service.CalcCubic(id)
	if err != nil {
		log.Printf("Error: %v", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var solution string

	if cubic < 10000 && cubic > 100 {
		solution = fmt.Sprintf("APTO. El valor cúbico %d se encuentra dentro de la escala subvencionable", cubic)
	} else {
		solution = fmt.Sprintf("NO APTO. El valor cúbico %d está fuera de la escala subvencionable", cubic)
	}

	ctx.JSON(nethttp.StatusOK, gin.H{"solution": solution})
}
