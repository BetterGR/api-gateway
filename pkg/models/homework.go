package models

import "time"

// Homework represents the current homework implementation
type Homework struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

// HomeworkSubmission represents the current homework submission implementation
type HomeworkSubmission struct {
	ID          string    `json:"id"`
	HomeworkID  string    `json:"homework_id" binding:"required"`
	StudentID   string    `json:"student_id" binding:"required"`
	Content     string    `json:"content" binding:"required"`
	SubmittedAt time.Time `json:"submitted_at"`
}

type getHomeworkRequest struct {
	CourseID string `uri:"courseId" binding:"required"`
}

type createHomeworkRequest struct {
	CourseID    string    `json:"course_id" binding:"required"`
	HomeworkID  string    `json:"homework_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date" binding:"required"`
}

type getHomeworkResponse struct {
	Homework Homework `json:"homework"`
}

type createHomeworkResponse struct {
	Homework Homework `json:"homework"`
}


