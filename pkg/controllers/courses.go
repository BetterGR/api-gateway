package controllers

import (
	"net/http"

	courseProtos "github.com/BetterGR/course-microservice/protos"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// InitCoursesGRPCClient initializes the course-microservice gRPC client connection.
func InitCoursesGRPCClient(address string) (courseProtos.CourseServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return courseProtos.NewCourseServiceClient(conn), nil
}

// GetAnnouncementHandler handles fetching an announcement for a course.
func GetAnnouncementHandler(c *gin.Context, client courseProtos.CourseServiceClient) {
	courseID := c.Param("courseId")

	klog.Infof("Fetching announcement for course: %s", courseID)

	req := &courseProtos.GetAnnouncementRequest{
		CourseId: courseID,
	}

	resp, err := client.GetAnnouncement(context.Background(), req)
	if err != nil {
		klog.Errorf("Failed to get announcement: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch announcement"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"course_id":    resp.CourseId,
		"announcement": resp.Announcement,
	})
}

func GetHomeworkHandler(c *gin.Context, client courseProtos.CourseServiceClient) {
    courseID := c.Param("courseId")
    homeworkID := c.Param("homeworkId")

    // Create the gRPC request
    req := &courseProtos.GetHomeworkRequest{
        CourseId:   courseID,
        HomeworkId: homeworkID,
    }

    // Call the gRPC GetHomework method
    resp, err := client.GetHomework(context.Background(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to fetch homework: %v", err),
        })
        return
    }

    // Return the homework details as JSON response
    c.JSON(http.StatusOK, gin.H{
        "course_id":     resp.CourseId,
        "homework_id":   resp.HomeworkId,
        "title":         resp.Title,
        "description":   resp.Description,
        "due_date":      resp.DueDate,
    })
}

func GetCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func CreateCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func UpdateCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func AddStudentToCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func RemoveStudentFromCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func AddStaffToCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func RemoveStaffFromCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func DeleteCourseHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func ListStudentsHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func ListStaffHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func AddHomeworkHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

func RemoveHomeworkHandler(c *gin.Context, grpcClient courseProtos.CourseServiceClient) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
}

