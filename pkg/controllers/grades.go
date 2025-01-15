package controllers

import (
	"context"
	"net/http"
	"time"

	gradesProtos "github.com/BetterGR/grades-microservice/protos"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/klog/v2"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// InitGradesGRPCClient initializes the grades-microservice gRPC client connection.
func InitGradesGRPCClient(address string) (gradesProtos.GradesServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return gradesProtos.NewGradesServiceClient(conn), nil
}

// GetStudentCourseGradesHandler handles REST requests and calls the gRPC Grades Microservice.
func GetStudentCourseGradesHandler(c *gin.Context, grpcClient gradesProtos.GradesServiceClient) {
	studentId := c.Param("studentId")
	courseId := c.Param("courseId")
	// Build gRPC request.
	request := &gradesProtos.GetStudentCourseGradesRequest{
		StudentId: studentId,
		CourseId:  courseId,
	}
	// Call the gRPC server.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := grpcClient.GetStudentCourseGrades(ctx, request)

	if err != nil {
		klog.Errorf("Error calling gRPC Grades Microservice: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch grades"})

		return
	}

	// Send response to the client.
	c.JSON(http.StatusOK, response.CourseGrades)
}

// GetStudentGradesHandler handles REST requests and calls the gRPC Grades Microservice to return
// all the student grades
func GetStudentGradesHandler(c *gin.Context, grpcClient gradesProtos.GradesServiceClient) {
	// Log all parameters for debugging
	klog.Infof("All params: %v", c.Params)

	studentId := c.Param("student_id")
	klog.Infof("Student ID from param: '%s'", studentId)

	// Build gRPC request.
	request := &gradesProtos.StudentId{StudentId: studentId}
	klog.Infof("Request built with student ID: '%s'", request.StudentId)

	// Call the gRPC server.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := grpcClient.GetStudentGrades(ctx, request)
	if err != nil {
		klog.Errorf("Error calling gRPC Grades Microservice: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch grades"})
		return
	}
	c.JSON(http.StatusOK, response)
}

func AddHomeworkGradeHandler(c *gin.Context, grpcClient gradesProtos.GradesServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func AddExamGradeHandler(c *gin.Context, grpcClient gradesProtos.GradesServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func UpdateHomeworkGradeHandler(c *gin.Context, grpcClient gradesProtos.GradesServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func UpdateExamGradeHandler(c *gin.Context, grpcClient gradesProtos.GradesServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func DeleteHomeworkGradeHandler(c *gin.Context, grpcClient gradesProtos.GradesServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func DeleteExamGradeHandler(c *gin.Context, grpcClient gradesProtos.GradesServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}
