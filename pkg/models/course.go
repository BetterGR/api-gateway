package models

// Course represents a course in the system
type Course struct {
	ID          string   `json:"id"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	TeacherID   string   `json:"teacher_id" binding:"required"`
	Students    []string `json:"students,omitempty"`
	IsActive    bool     `json:"is_active"`
}

// CourseEnrollment represents a student's enrollment in a course
type CourseEnrollment struct {
	CourseID  string `json:"course_id" binding:"required"`
	StudentID string `json:"student_id" binding:"required"`
}
