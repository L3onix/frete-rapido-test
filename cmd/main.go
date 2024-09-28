package main

import (
	"github.com/L3onix/frete-rapido-test/internal/router"
	"github.com/L3onix/frete-rapido-test/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	dbConnection, err := db.ConnectionDB()
	if err != nil {
		panic(err)
	}

	server := router.SetupRoutes(dbConnection, gin.Default())
	server.Run(":8080")
}
