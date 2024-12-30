package main

import (
	"api-gateway/controllers"
	"api-gateway/middleware"
	"fmt"
	"os"

	"k8s.io/klog/v2"

	"github.com/gin-gonic/gin"
)

const (
	address = "localhost:50051"
)

func main() {
	klog.InitFlags(nil)
	defer klog.Flush()

	// Initialize the gRPC client connection.
	grpcClient, err := controllers.InitGRPCClient(address)
	if err != nil {
		klog.Fatalf("Failed to initialize gRPC client, %v", err)
	}
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	// Rest endpoints.
	router.POST("/api/login", middleware.LoginHandler)
	router.GET("/api/dashboard", controllers.GetDashboardData)
	router.GET("/api/grades/:student_id/:courseId", func(c *gin.Context) {
		controllers.GetStudentGradesHandler(c, grpcClient)
	})

	// Get the port from the environment variable, default to 1234 if not set
	port := os.Getenv("API_GATEWAY_PORT")
	if port == "" {
		klog.Fatalf("API_GATEWAY_PORT is not set")
	}
	// Start the server on the specified port
	router.Run(fmt.Sprintf(":%s", port))
}
