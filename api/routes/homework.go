package routes

import (
	"os"

	"github.com/BetterGR/api-gateway/pkg/controllers"
	homeworkProtos "github.com/BetterGR/homework-microservice/protos"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

// InitiateHomeWorkMicroservice initialize homework microservice
func InitiateHomeWorkMicroservice(router *gin.Engine) {
	_, err := RegisterHomeWorkRoutes(router)
	if err != nil {
		klog.Fatalf("Failed to register homework routes, %v", err)
	}
}

func RegisterHomeWorkRoutes(router *gin.Engine) (homeworkProtos.HomeworkServiceClient, error) {
	// Initialize the gRPC client connection.
	homeworkAddress := os.Getenv("HOMEWORK_ADDRESS")
	klog.Infof("Homework address: %s", homeworkAddress)
	grpcClient, err := controllers.InitHomeWorkGRPCClient(homeworkAddress)
	if err != nil {
		klog.Fatalf("Failed to initialize gRPC client, %v", err)
	}

	// Rest endpoints.
	router.GET("/api/homework/:courseId", func(c *gin.Context) {
		controllers.GetHomeworkHandler(c, grpcClient)
	})
	router.POST("/api/homework", func(c *gin.Context) {
		controllers.CreateHomeworkHandler(c, grpcClient)
	})

	return grpcClient, nil
}
