package routes

import (
	"github.com/BetterGR/api-gateway/pkg/controllers"
	"github.com/BetterGR/api-gateway/pkg/middleware"
	"github.com/gin-gonic/gin"
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

// RegisterUIServices registers the UI services.
func RegisterUIServices(router *gin.Engine) {
	router.POST("/api/login", controllers.LoginHandler)
	router.GET("/api/dashboard", controllers.GetDashboardData)
}

// InitiateRoutes initializes the routes of Microservices and UI services for the API Gateway.
func InitiateRoutes(router *gin.Engine) {
	router.Use(middleware.CORSMiddleware())
	klog.Info("Registering routes")
	RegisterUIServices(router)
	RegisterMicroservices(router)
}
