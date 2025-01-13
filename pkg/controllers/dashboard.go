package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDashboardData(c *gin.Context) {
	data := gin.H{
		"headers": []gin.H{
			{"title": "Student Profile", "content": "Details about the student"},
			{"title": "Courses", "content": "List of courses"},
			{"title": "Grades", "content": "Student grades"},
			{"title": "Tips", "content": "Useful academic tips"},
		},
	}

	c.JSON(http.StatusOK, data)
}
