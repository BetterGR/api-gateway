package models

import "time"

// Staff represents a staff member in the system
type Staff struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Role      string    `json:"role" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Courses   []Course  `json:"courses,omitempty"`
}

// New structs for requests and responses

type getStaffMemberRequest struct {
	ID string `uri:"staffId" binding:"required"`
}

type createStaffRequest struct {
	ID        string `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

type updateStaffRequest struct {
	ID        string `uri:"staffId" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

type deleteStaffMemberRequest struct {
	ID string `uri:"staffId" binding:"required"`
}

type getStaffMemberResponse struct {
	Staff Staff `json:"staff"`
}

type createStaffMemberRequest struct {
	Staff Staff `json:"staff"`
}

type updateStaffMemberRequest struct {
	Staff Staff `json:"staff"`
}

type updateStaffMemberResponse struct {
	Staff Staff `json:"staff"`
}

type deleteStaffMemberResponse struct {
}

type getStaffCoursesRequest struct {
	ID string `uri:"staffId" binding:"required"`
}

type getStaffCoursesResponse struct {
	Courses []Course `json:"courses"`
}

type createStaffMemberResponse struct {
	Staff Staff `json:"staff"`
}

