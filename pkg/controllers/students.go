package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/BetterGR/students-microservice/protos"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/klog/v2"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// InitStudentsGRPCClient initializes the students-microservice gRPC client connection.
func InitStudentsGRPCClient(address string) (protos.StudentsServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return protos.NewStudentsServiceClient(conn), nil
}

// GetStudentCourssHandler handles REST requests and calls the gRPC Students Microservice.
func GetStudentCoursesHandler(c *gin.Context, grpcClient protos.StudentsServiceClient) {
	klog.Info("Sabry")
	studentId := c.Param("studentId")
	// Build gRPC request.
	request := &protos.GetStudentCoursesRequest{
		Id: studentId,
	}
	// Call the gRPC server.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := grpcClient.GetStudentCourses(ctx, request)

	if err != nil {
		klog.Errorf("Error calling gRPC Students Microservice: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})

		return
	}

	// Send response to the client.
	c.JSON(http.StatusOK, response.GetCourses())
}
