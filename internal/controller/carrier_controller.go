package controller

import (
	"net/http"
	"strconv"

	"github.com/L3onix/frete-rapido-test/internal/model"
	"github.com/L3onix/frete-rapido-test/internal/usecase"
	"github.com/gin-gonic/gin"
)

type carrierController struct {
	useCase usecase.CarrierUseCase
}

func NewCarrierController(useCase usecase.CarrierUseCase) carrierController {
	return carrierController{
		useCase: useCase,
	}
}

func (c *carrierController) GetCarriers(ctx *gin.Context) {
	carriers, err := c.useCase.GetCarriers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, carriers)
}

func (c *carrierController) CreateCarrier(ctx *gin.Context) {
	var carrier model.Carrier
	err := ctx.ShouldBindJSON(&carrier)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCarrier, err := c.useCase.CreateCarrier(carrier)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusCreated, createdCarrier)
}

func (c *carrierController) GetCarrierById(ctx *gin.Context) {
	id := ctx.Param("carrierId")
	if id == "" {
		response := model.Response{
			Message: "ID não pode ser nulo!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	carrierId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID precisa ser um número!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	carrier, err := c.useCase.GetCarrierById(carrierId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if carrier == nil {
		response := model.Response{
			Message: "Carrier não encontrado.",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, carrier)
}
