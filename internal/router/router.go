package router

import (
	"database/sql"

	"github.com/L3onix/frete-rapido-test/internal/controller"
	"github.com/L3onix/frete-rapido-test/internal/repository"
	"github.com/L3onix/frete-rapido-test/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(database *sql.DB, server *gin.Engine) *gin.Engine {
	CarrierRepository := repository.NewCarrierRepository(database)
	CarrierUseCase := usecase.NewCarrierUseCase(CarrierRepository)
	CarrierController := controller.NewCarrierController(CarrierUseCase)

	QuoteController := controller.NewQuoteController()

	MetricUsecase := usecase.NewMetricUseCase(CarrierUseCase)
	MetricController := controller.NewMetricController(MetricUsecase)

	server.GET("/carrier", CarrierController.GetCarriers)
	server.GET("/carrier/:carrierId", CarrierController.GetCarrierById)
	server.POST("/carrier", CarrierController.CreateCarrier)

	// rotas definidas pelo teste
	server.GET("/metrics", MetricController.GetMetrics)
	server.POST("/quote", QuoteController.CreateQuote)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	return server
}
