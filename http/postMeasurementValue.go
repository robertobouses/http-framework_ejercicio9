package http

/*
import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

func (h *Http) PostMeasurementValue(ctx *gin.Context) {
	var measurementValueRequest entity.MeasurementValueRequest
	if err := ctx.ShouldBindJSON(&measurementValueRequest); err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error al hacer ShouldBindJSON": err.Error()})
		return
	}

	cubic, err := h.service.CalcCubic(measurementValueRequest.ID)
	amount=

}
*/
