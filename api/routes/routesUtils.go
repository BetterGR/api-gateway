package routes

import (
	"github.com/BetterGR/api-gateway/pkg/controllers"
	"github.com/BetterGR/api-gateway/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"k8s.io/klog/v2"
)

// RegisterMicroservices registers the microservices.
func RegisterMicroservices(router *gin.Engine) {
	InitiateGradesMicroservice(router)
	InitiateStudentsMicroservice(router)
	InitiateCoursesMicroservice(router)
	InitiateHomeWorkMicroservice(router)
	InitiateStaffMicroservice(router)
}

// InitiateRoutes initializes all routes for the API Gateway
func InitiateRoutes(router *gin.Engine) {
	router.Use(middleware.CORSMiddleware())

	// Setup Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register callback endpoint
	router.POST("/api/callback", controllers.HandleCallback)

	klog.Info("Registering routes")
	RegisterMicroservices(router)
}
