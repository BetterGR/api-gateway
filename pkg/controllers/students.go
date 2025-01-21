package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	studentsProtos "github.com/BetterGR/students-microservice/protos"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/klog/v2"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// InitStudentsGRPCClient initializes the students-microservice gRPC client connection.
func InitStudentsGRPCClient(address string) (studentsProtos.StudentsServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return studentsProtos.NewStudentsServiceClient(conn), nil
}

// GetStudentCourssHandler handles REST requests and calls the gRPC Students Microservice.
func GetStudentCoursesHandler(c *gin.Context, grpcClient studentsProtos.StudentsServiceClient) {
	// Extract token from Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token provided"})
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")

	studentId := c.Param("studentId")
	// Build gRPC request with token
	request := &studentsProtos.GetStudentCoursesRequest{
		Token: token, // Add token as the first field
		Id:    studentId,
	}
	// Call the gRPC server.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	klog.Infof("making request for student courses")
	response, err := grpcClient.GetStudentCourses(ctx, request)

	if err != nil {
		klog.Errorf("Error calling gRPC Students Microservice: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})

		return
	}

	// Send response to the client.
	c.JSON(http.StatusOK, response.GetCourses())
}

func CreateStudentHandler(c *gin.Context, grpcClient studentsProtos.StudentsServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func GetStudentHandler(c *gin.Context, grpcClient studentsProtos.StudentsServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func UpdateStudentHandler(c *gin.Context, grpcClient studentsProtos.StudentsServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func GetStudentGradesHandlerStudent(c *gin.Context, grpcClient studentsProtos.StudentsServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func DeleteStudentHandler(c *gin.Context, grpcClient studentsProtos.StudentsServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}
