package controller

import (
	"net/http"

	"github.com/L3onix/frete-rapido-test/external/frete_rapido/service"
	"github.com/L3onix/frete-rapido-test/internal/model"

	// "github.com/L3onix/frete-rapido-test/internal/usecase"
	"github.com/gin-gonic/gin"
)

type quoteController struct {
}

func NewQuoteController() quoteController {
	return quoteController{}
}

func (q *quoteController) CreateQuote(ctx *gin.Context) {
	var quote model.Quote
	err := ctx.ShouldBindJSON(&quote)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	simulateResult, err := service.FetchQuoteSimulate(quote)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// qu := usecase.NewQuoteUseCase()
	// qu.StoreResult(simulateResult)

	ctx.JSON(http.StatusOK, simulateResult)
	return
}
