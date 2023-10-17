package http

import (
	nethttp "net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Http) PostInvoiced(ctx *gin.Context, amountStr string) {
	amount64, err := strconv.ParseFloat(amountStr, 32)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, err)
		return
	}
	amount := float32(amount64)
	value, err := h.service.ChangeInvoiced(amount)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"Error:": err.Error()})

	}
	ctx.JSON(nethttp.StatusOK, value)
}
