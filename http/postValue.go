package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

func (h *Http) PostValue(ctx *gin.Context) {
	var newValue entity.Value
	err := ctx.BindJSON(&newValue)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, err)
		return
	}
	err = h.service.CreateValue(newValue)
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error llamada http a app": err.Error()})
		return
	}
	if newValue.Amount > 10 {
		ctx.JSON(200, "el valor sobrepasa 10")
	}
	if newValue.Amount <= 10 {
		ctx.JSON(200, "el valor es inferior o igual a 10")
	}
	ctx.JSON(200, newValue)
}
