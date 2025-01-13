package main

import (
	roterApi "github.com/BetterGR/api-gateway/api/routes"
	"os"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

const (
	address = "localhost:50051"
)

func main() {
	klog.InitFlags(nil)
	defer klog.Flush()
	// Get the port from the environment variable, default to 1234 if not set
	port := os.Getenv("API_GATEWAY_PORT")
	if port == "" {
		klog.Fatalf("API_GATEWAY_PORT is not set")
	}
	router := gin.New()
	roterApi.InitiateRoutes(router)
	err := router.Run(":" + port)
	if err != nil {
		klog.Fatalf("Failed to start the server, %v", err)
	}
}
