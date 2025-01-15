package routes

import (
	"os"

	"github.com/BetterGR/api-gateway/pkg/controllers"
	staffProtos "github.com/BetterGR/staff-microservice/protos"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

// InitiateStaffMicroservice initialize staff microservice
func InitiateStaffMicroservice(router *gin.Engine) {
	_, err := RegisterStaffRoutes(router)
	if (err != nil) {
		klog.Fatalf("Failed to register staff routes, %v", err)
	}
}

func RegisterStaffRoutes(router *gin.Engine) (staffProtos.StaffServiceClient, error) {
	// Initialize the gRPC client connection.
	staffAddress := os.Getenv("STAFF_ADDRESS")
	klog.Infof("Staff address: %s", staffAddress)
	grpcClient, err := controllers.InitStaffGRPCClient(staffAddress)
	if err != nil {
		klog.Fatalf("Failed to initialize gRPC client, %v", err)
	}

	// Rest endpoints.
	router.GET("/api/staff/:staffId", func(c *gin.Context) {
		controllers.GetStaffMemberHandler(c, grpcClient)
	})
	router.GET("/api/staff/:staffId/courses", func(c *gin.Context) {
		controllers.GetCoursesListHandler(c, grpcClient)
	})
	router.POST("/api/staff", func(c *gin.Context) {
		controllers.CreateStaffMemberHandler(c, grpcClient)
	})
	router.PUT("/api/staff/:staffId", func(c *gin.Context) {
		controllers.UpdateStaffMemberHandler(c, grpcClient)
	})
	router.DELETE("/api/staff/:staffId", func(c *gin.Context) {
		controllers.DeleteStaffMemberHandler(c, grpcClient)
	})

	return grpcClient, nil
}
