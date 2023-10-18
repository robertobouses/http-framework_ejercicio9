package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

func (h *Http) PostNameValue(ctx *gin.Context) {
	var nameValueRequest entity.NameValueRequest
	if err := ctx.ShouldBindJSON(&nameValueRequest); err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error ShouldBind": err})
		return
	}

	names, err := h.service.SearchNameValue(nameValueRequest)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, err)
	}
	ctx.JSON(200, names)

}
