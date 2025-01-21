package models

// Course represents a course in the system
type Course struct {
	ID          string   `json:"id"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	TeacherID   string   `json:"teacherId" binding:"required"`
	Students    []string `json:"students,omitempty"`
	IsActive    bool     `json:"isActive"`
}

// CourseEnrollment represents a student's enrollment in a course
type CourseEnrollment struct {
	CourseID  string `json:"courseId" binding:"required"`
	StudentID string `json:"studentId" binding:"required"`
}

type getCourseRequest struct {
	ID string `uri:"courseId" binding:"required"`
}

type createCourseRequest struct {
	ID          string   `json:"id" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	TeacherID   string   `json:"teacherId" binding:"required"`
	IsActive    bool     `json:"isActive"`
}

type updateCourseRequest struct {
	ID          string   `uri:"courseId" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	TeacherID   string   `json:"teacherId" binding:"required"`
	IsActive    bool     `json:"isActive"`
}

type addStudentToCourseRequest struct {
	ID        string `uri:"courseId" binding:"required"`
	StudentID string `json:"studentId" binding:"required"`
}

type removeStudentFromCourseRequest struct {
	ID        string `uri:"courseId" binding:"required"`
	StudentID string `uri:"studentId" binding:"required"`
}

type addStaffToCourseRequest struct {
	ID      string `uri:"courseId" binding:"required"`
	StaffID string `json:"staffId" binding:"required"`
}

type removeStaffFromCourseRequest struct {
	ID      string `uri:"courseId" binding:"required"`
	StaffID string `uri:"staffId" binding:"required"`
}

type deleteCourseRequest struct {
	ID string `uri:"courseId" binding:"required"`
}

type getCourseStudentsRequest struct {
	ID string `uri:"courseId" binding:"required"`
}

type getCourseStaffRequest struct {
	ID string `uri:"courseId" binding:"required"`
}

type getCourseResponse struct {
	Course Course `json:"course"`
}

type createCourseResponse struct {
	Course Course `json:"course"`
}

type updateCourseResponse struct {
	Course Course `json:"course"`
}

type addStudentToCourseResponse struct {
}

type removeStudentFromCourseResponse struct {
}

type addStaffToCourseResponse struct {
}

type removeStaffFromCourseResponse struct {
}

type deleteCourseResponse struct {
}

type getCourseStudentsResponse struct {
	Students []string `json:"students"`
}

type getCourseStaffResponse struct {
	Staff []string `json:"staff"`
}

