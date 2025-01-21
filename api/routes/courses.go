package routes

import (
	"os"

	"github.com/BetterGR/api-gateway/pkg/controllers"
	courseProtos "github.com/BetterGR/course-microservice/protos"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

// InitiateCoursesMicroservice initialize courses microservice
func InitiateCoursesMicroservice(router *gin.Engine) {
	_, err := RegisterCoursesRoutes(router)
	if err != nil {
		klog.Fatalf("Failed to register courses routes, %v", err)
	}
}

func RegisterCoursesRoutes(router *gin.Engine) (courseProtos.CourseServiceClient, error) {
	// Initialize the gRPC client connection.
	coursesAddress := os.Getenv("COURSES_ADDRESS")
	klog.Infof("Courses address: %s", coursesAddress)
	grpcClient, err := controllers.InitCoursesGRPCClient(coursesAddress)
	if err != nil {
		klog.Fatalf("Failed to initialize gRPC client, %v", err)
	}

	// Rest endpoints.
	router.GET("/api/courses/:courseId", func(c *gin.Context) {
		controllers.GetCourseHandler(c, grpcClient)
	})

	router.GET("/api/courses/:courseId/announcement", func(c *gin.Context) {
		controllers.GetAnnouncementHandler(c, grpcClient)
	})

	router.POST("/api/courses", func(c *gin.Context) {
		controllers.CreateCourseHandler(c, grpcClient)
	})
	router.PUT("/api/courses/:courseId", func(c *gin.Context) {
		controllers.UpdateCourseHandler(c, grpcClient)
	})
	router.POST("/api/courses/:courseId/students", func(c *gin.Context) {
		controllers.AddStudentToCourseHandler(c, grpcClient)
	})
	router.DELETE("/api/courses/:courseId/students/:studentId", func(c *gin.Context) {
		controllers.RemoveStudentFromCourseHandler(c, grpcClient)
	})
	router.POST("/api/courses/:courseId/staff", func(c *gin.Context) {
		controllers.AddStaffToCourseHandler(c, grpcClient)
	})
	router.DELETE("/api/courses/:courseId/staff/:staffId", func(c *gin.Context) {
		controllers.RemoveStaffFromCourseHandler(c, grpcClient)
	})
	router.DELETE("/api/courses/:courseId", func(c *gin.Context) {
		controllers.DeleteCourseHandler(c, grpcClient)
	})
	router.GET("/api/courses/:courseId/students", func(c *gin.Context) {
		controllers.ListStudentsHandler(c, grpcClient)
	})
	router.GET("/api/courses/:courseId/staff", func(c *gin.Context) {
		controllers.ListStaffHandler(c, grpcClient)
	})
	router.POST("/api/courses/:courseId/homework", func(c *gin.Context) {
		controllers.AddHomeworkHandler(c, grpcClient)
	})
	//router.GET("/api/courses/:courseId/homework/:homeworkId", func(c *gin.Context) {
	//	controllers.GetHomeworkHandler(c, grpcClient)
	//})
	router.DELETE("/api/courses/:courseId/homework/:homeworkId", func(c *gin.Context) {
		controllers.RemoveHomeworkHandler(c, grpcClient)
	})

	return grpcClient, nil
}
