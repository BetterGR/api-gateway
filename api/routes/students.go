package routes

import (
	"os"

	"github.com/BetterGR/api-gateway/pkg/controllers"
	studentsProtos "github.com/BetterGR/students-microservice/protos"
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

func RegisterStudentsRoutes(router *gin.Engine) (studentsProtos.StudentsServiceClient, error) {
	// Initialize the gRPC client connection.
	studentsAddress := os.Getenv("STUDENTS_ADDRESS")
	klog.Infof("Students address: %s", studentsAddress)
	grpcClient, err := controllers.InitStudentsGRPCClient(studentsAddress)
	if err != nil {
		klog.Fatalf("Failed to initialize gRPC client, %v", err)
	}

	// Rest endpoints.
	router.GET("/api/students/:studentId/courses", func(c *gin.Context) {
		controllers.GetStudentCoursesHandler(c, grpcClient)
	})
	router.POST("/api/students", func(c *gin.Context) {
		controllers.CreateStudentHandler(c, grpcClient)
	})
	router.GET("/api/students/:studentId", func(c *gin.Context) {
		controllers.GetStudentHandler(c, grpcClient)
	})
	router.PUT("/api/students/:studentId", func(c *gin.Context) {
		controllers.UpdateStudentHandler(c, grpcClient)
	})
	router.GET("/api/students/:studentId/grades", func(c *gin.Context) {
		controllers.GetStudentGradesHandlerStudent(c, grpcClient)
	})
	router.DELETE("/api/students/:studentId", func(c *gin.Context) {
		controllers.DeleteStudentHandler(c, grpcClient)
	})

	return grpcClient, nil
}
