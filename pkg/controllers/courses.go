package controllers

import (
	"net/http"

	courseProtos "github.com/BetterGR/course-microservice/protos"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// InitCoursesGRPCClient initializes the course-microservice gRPC client connection.
func InitCoursesGRPCClient(address string) (courseProtos.CourseServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return courseProtos.NewCourseServiceClient(conn), nil
}

func GetCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func CreateCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func UpdateCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func AddStudentToCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func RemoveStudentFromCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func AddStaffToCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func RemoveStaffFromCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func DeleteCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func ListStudentsHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func ListStaffHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func AddHomeworkHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func RemoveHomeworkHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

