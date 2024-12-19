package main

import (
	"api-gateway/controllers"
	"api-gateway/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the gRPC client connection.
	controllers.InitGRPCClient()

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	// Rest endpoints.
	router.POST("/api/login", controllers.LoginHandler)
	router.GET("/api/dashboard", controllers.GetDashboardData)
	router.GET("/api/grades/:student_id/:courseId", controllers.GetStudentGradesHandler)

	router.Run(":8080")
}
