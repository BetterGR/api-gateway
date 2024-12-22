package controllers

import (
	"context"
	"github.com/BetterGR/grades-microservice/protos"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitGRPCClient(address string) (protos.GradesServiceClient, error) {
	// Initialize the gRPC client connection.
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return protos.NewGradesServiceClient(conn), nil
}

// GetStudentGradesHandler handles REST requests and calls the gRPC Grades Microservice.
func GetStudentGradesHandler(c *gin.Context, grpcClient protos.GradesServiceClient) {
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
	if err != nil {
		log.Printf("Error calling gRPC Grades Microservice: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch grades"})

		return
	}

	// Send response to the client.
	c.JSON(http.StatusOK, response.CourseGrades)
}
