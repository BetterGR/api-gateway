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

// Note: Unimplemented features

type HomeworkAttachment struct {
	NotImplementedYet
}

type HomeworkFeedback struct {
	NotImplementedYet
}

type HomeworkBulkGrade struct {
	NotImplementedYet
}
