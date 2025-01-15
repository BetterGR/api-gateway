package main

import (
	"github.com/BetterGR/api-gateway/api/routes"
	_ "github.com/BetterGR/api-gateway/docs" // Required for swagger docs
	"github.com/gin-gonic/gin"
)

// @title BetterGR API Gateway
// @version 1.0
// @description API Gateway for the BetterGR grading system
// @host localhost:1234
// @BasePath /api
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":1234")
}
