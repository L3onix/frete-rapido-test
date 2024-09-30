package controller

import (
	"net/http"
	"strconv"

	"github.com/L3onix/frete-rapido-test/internal/usecase"
	"github.com/gin-gonic/gin"
)

type metricController struct {
	usecase usecase.MetricUsecase
}

func NewMetricController(metricUsecase usecase.MetricUsecase) metricController {
	return metricController{
		usecase: metricUsecase,
	}
}

func (mc *metricController) GetMetrics(ctx *gin.Context) {
	lastQuotes := ctx.Query("last_quotes")
	if lastQuotes != "" {
		if _, err := strconv.Atoi(lastQuotes); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "last_quotes precisa ser um número!",
			})
			return
		}
	}
	dispatcherID := ctx.Query("dispatcher_id")

	metrics, err := mc.usecase.GetMetrics(lastQuotes, dispatcherID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if len(metrics.Carriers) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Banco de dados não possui métricas",
		})
		return
	}

	ctx.JSON(http.StatusOK, metrics)
	return
}
