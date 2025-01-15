package routes

import (
	"os"

	"github.com/BetterGR/api-gateway/pkg/controllers"
	"github.com/BetterGR/students-microservice/protos"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

// InitiateStudentsMicroservice initialize students microservice
func InitiateStudentsMicroservice(router *gin.Engine) {
	_, err := RegisterStudentsRoutes(router)
	if err != nil {
		klog.Fatalf("Failed to register students routes, %v", err)
	}
}

func RegisterStudentsRoutes(router *gin.Engine) (protos.StudentsServiceClient, error) {
	// Initialize the gRPC client connection.
	studentsAddress := os.Getenv("STUDENTS_ADDRESS")
	klog.Infof("Students address: %s", studentsAddress)
	grpcClient, err := controllers.InitStudentsGRPCClient(studentsAddress)
	if err != nil {
		klog.Fatalf("Failed to initialize gRPC client, %v", err)
	}

	// Rest endpoints.
	router.GET("/api/students/:studentId", func(c *gin.Context) {
		controllers.GetStudentCoursesHandler(c, grpcClient)
	})
	return grpcClient, nil
}
