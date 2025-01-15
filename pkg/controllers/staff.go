package controllers

import (
	"net/http"

	staffProtos "github.com/BetterGR/staff-microservice/protos"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// InitStaffGRPCClient initializes the staff-microservice gRPC client connection.
func InitStaffGRPCClient(address string) (staffProtos.StaffServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return staffProtos.NewStaffServiceClient(conn), nil
}

func GetStaffMemberHandler(c *gin.Context, grpcClient staffProtos.StaffServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func GetCoursesListHandler(c *gin.Context, grpcClient staffProtos.StaffServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func CreateStaffMemberHandler(c *gin.Context, grpcClient staffProtos.StaffServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func UpdateStaffMemberHandler(c *gin.Context, grpcClient staffProtos.StaffServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func DeleteStaffMemberHandler(c *gin.Context, grpcClient staffProtos.StaffServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

