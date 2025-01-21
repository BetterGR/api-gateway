package models

import "time"

type Student struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	Courses   []Course  `json:"courses,omitempty"`
}

type getStudentRequest struct {
	ID string `uri:"studentId" binding:"required"`
}

type createStudentRequest struct {
	ID 	  string `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

type updateStudentRequest struct {
	ID        string `uri:"studentId" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

type deleteStudentRequest struct {
	ID string `uri:"studentId" binding:"required"`
}

type getStudentCoursesRequest struct {
	ID string `uri:"studentId" binding:"required"`
}

type getStudentCourseGradesRequest struct {
	StudentID string `uri:"studentId" binding:"required"`
	CourseID  string `uri:"courseId" binding:"required"`
}

type getStudentGradesRequest struct {
	ID string `uri:"studentId" binding:"required"`
}

type getStudentResponse struct {
	Student Student `json:"student"`
}

type createStudentResponse struct {
	Student Student `json:"student"`
}

type updateStudentResponse struct {
	Student Student `json:"student"`
}

type deleteStudentResponse struct {
}

type getStudentCoursesResponse struct {
	Courses []Course `json:"courses"`
}

type getStudentCourseGradesResponse struct {
	CourseID  string `json:"course_id"`
    Grades    []Grade `json:"grades"`
}

type getStudentGradesResponse struct{
	Grades    []Grade `json:"grades"`
}
