package models

import "time"

// Grade represents a student's grade as implemented in the grades controller
type Grade struct {
	ID        string    `json:"id"`
	StudentID string    `json:"student_id" binding:"required"`
	CourseID  string    `json:"course_id" binding:"required"`
	Score     float64   `json:"score" binding:"required"`
	GradedAt  time.Time `json:"graded_at"`
	Comments  string    `json:"comments"`
}

// GradeSubmission represents the current implementation of grade submission
type GradeSubmission struct {
	StudentID string  `json:"student_id" binding:"required"`
	CourseID  string  `json:"course_id" binding:"required"`
	Score     float64 `json:"score" binding:"required"`
	Comments  string  `json:"comments"`
}

// Note: Unimplemented features

type GradeBulkSubmission struct {
	NotImplementedYet
}

type GradeAnalytics struct {
	NotImplementedYet
}

type GradeExport struct {
	NotImplementedYet
}
