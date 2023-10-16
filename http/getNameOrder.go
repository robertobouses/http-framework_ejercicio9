package http

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h *Http) GetNameOrder(ctx *gin.Context) {
	listaOrdenadaNombres, err := h.service.PrintNameOrder()
	if err != nil {
		log.Printf("Error al obtener la lista ordenada", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.Print("Listado ordenado de nombres de importe mayor a menor:", listaOrdenadaNombres)
	ctx.JSON(nethttp.StatusOK, listaOrdenadaNombres)
}
