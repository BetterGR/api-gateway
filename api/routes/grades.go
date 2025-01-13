package routes

import (
	"github.com/BetterGR/api-gateway/pkg/controllers"
	"github.com/BetterGR/grades-microservice/protos"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

const (
	address = "localhost:50051"
)

func RegisterGradesRoutes(router *gin.Engine) (protos.GradesServiceClient, error) {
	// Initialize the gRPC client connection.
	grpcClient, err := controllers.InitGradesGRPCClient(address)
	if err != nil {
		klog.Fatalf("Failed to initialize gRPC client, %v", err)
	}

	// Rest endpoints.
	router.GET("/api/grades/:student_id/:courseId", func(c *gin.Context) {
		controllers.GetStudentGradesHandler(c, grpcClient)
	})

	return grpcClient, nil
}
