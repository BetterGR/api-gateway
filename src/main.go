package main

import (
	"api-gateway/controllers"
	"api-gateway/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	address = "localhost:50051"
)

func main() {
	// Initialize the gRPC client connection.
	grpcClient, err := controllers.InitGRPCClient(address)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	// Rest endpoints.
	router.POST("/api/login", controllers.LoginHandler)
	router.GET("/api/dashboard", controllers.GetDashboardData)
	router.GET("/api/grades/:student_id/:courseId", func(c *gin.Context) {
		controllers.GetStudentGradesHandler(c, grpcClient)
	})

	router.Run(":8080")
}
