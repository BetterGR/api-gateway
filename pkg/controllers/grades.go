package controllers

import (
	"context"
	"github.com/BetterGR/grades-microservice/protos"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/klog/v2"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// InitGradesGRPCClient initializes the grades-microservice gRPC client connection.
func InitGradesGRPCClient(address string) (protos.GradesServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return protos.NewGradesServiceClient(conn), nil
}

// GetStudentCourseGradesHandler handles REST requests and calls the gRPC Grades Microservice.
func GetStudentCourseGradesHandler(c *gin.Context, grpcClient protos.GradesServiceClient) {
	studentId := c.Param("studentId")
	courseId := c.Param("courseId")
	// Build gRPC request.
	request := &protos.GetStudentCourseGradesRequest{
		StudentId: studentId,
		CourseId:  courseId,
	}
	// Call the gRPC server.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := grpcClient.GetStudentCourseGrades(ctx, request)
	logger := klog.FromContext(ctx)
	if err != nil {
		logger.Info("Error calling gRPC Grades Microservice: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch grades"})

		return
	}

	// Send response to the client.
	c.JSON(http.StatusOK, response.CourseGrades)
}

// GetStudentGradesHandler handles REST requests and calls the gRPC Grades Microservice to return
// all the student grades
func GetStudentGradesHandler(c *gin.Context, grpcClient protos.GradesServiceClient) {
	studentId := c.Param("studentId")
	klog.Info("Student ID: %s", studentId)
	// Build gRPC request.
	request := &protos.StudentId{StudentId: studentId}

	// Call the gRPC server.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := grpcClient.GetStudentGrades(ctx, request)
	logger := klog.FromContext(ctx)
	if err != nil {
		logger.Info("Error calling gRPC Grades Microservice: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch grades"})

		return
	}
	c.JSON(http.StatusOK, response)

}
