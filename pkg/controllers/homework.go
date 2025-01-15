package controllers

import (
	"net/http"

	homeworkProtos "github.com/BetterGR/homework-microservice/protos"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// InitHomeWorkGRPCClient initializes the homework-microservice gRPC client connection.
func InitHomeWorkGRPCClient(address string) (homeworkProtos.HomeworkServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return homeworkProtos.NewHomeworkServiceClient(conn), nil
}

func GetHomeworkHandler(c *gin.Context, grpcClient homeworkProtos.HomeworkServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func CreateHomeworkHandler(c *gin.Context, grpcClient homeworkProtos.HomeworkServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

