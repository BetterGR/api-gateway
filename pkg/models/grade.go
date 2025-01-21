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

type addHomeworkGradeRequest struct {
	GradeSubmission GradeSubmission `json:"grade_submission" binding:"required"`
	HomeworkID       string          `json:"homework_id" binding:"required"`
}

type addExamGradeRequest struct {
	GradeSubmission GradeSubmission `json:"grade_submission" binding:"required"`
	ExamID           string          `json:"exam_id" binding:"required"`
}

type updateHomeworkGradeRequest struct {
	GradeSubmission GradeSubmission `json:"grade_submission" binding:"required"`
	HomeworkID       string          `json:"homework_id" binding:"required"`
}

type updateExamGradeRequest struct {
	GradeSubmission GradeSubmission `json:"grade_submission" binding:"required"`
	ExamID           string          `json:"exam_id" binding:"required"`
}

type deleteHomeworkGradeRequest struct {
	StudentID string `json:"student_id" binding:"required"`
	CourseID  string `json:"course_id" binding:"required"`
	HomeworkID string `json:"homework_id" binding:"required"`
}

type deleteExamGradeRequest struct {
	StudentID string `json:"student_id" binding:"required"`
	CourseID  string `json:"course_id" binding:"required"`
	ExamID    string `json:"exam_id" binding:"required"`
}

type getHomeworkGradeRequest struct {
	StudentID string `json:"student_id" binding:"required"`
	CourseID  string `json:"course_id" binding:"required"`
	HomeworkID string `json:"homework_id" binding:"required"`
}

type getExamGradeRequest struct {
	StudentID string `json:"student_id" binding:"required"`
	CourseID  string `json:"course_id" binding:"required"`
	ExamID    string `json:"exam_id" binding:"required"`
}

type getStudentCourseGradeRequest struct {
	StudentID string `json:"student_id" binding:"required"`
	CourseID  string `json:"course_id" binding:"required"`
}

type getHomeworkGradesRespo struct {
	StudentID string `json:"student_id" binding:"required"`
	CourseID  string `json:"course_id" binding:"required"`
	HomeworkID string `json:"homework_id" binding:"required"`
}

type getExamGrades struct {
	StudentID string `json:"student_id" binding:"required"`
	CourseID  string `json:"course_id" binding:"required"`
	ExamID    string `json:"exam_id" binding:"required"`
}

type getStudentCourseGrades struct {
	StudentID string `json:"student_id" binding:"required"`
	CourseID  string `json:"course_id" binding:"required"`
}

type addHomeworkGradeResponse struct {
	Grade Grade `json:"grade"`
}

type addExamGradeResponse struct {
	Grade Grade `json:"grade"`
}

type updateHomeworkGradeResponse struct {
	Grade Grade `json:"grade"`
}

type updateExamGradeResponse struct {
	Grade Grade `json:"grade"`
}

type deleteHomeworkGradeResponse struct {
}

type deleteExamGradeResponse struct {
}

type getHomeworkGradeResponse struct {
	Grade Grade `json:"grade"`
}

type getExamGradeResponse struct {
	Grade Grade `json:"grade"`
}

type getStudentCourseGradeResponse struct {
	Grade Grade `json:"grade"`
}
