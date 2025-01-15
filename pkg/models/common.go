package models

// NotImplementedYet represents a response for features not yet implemented
type NotImplementedYet struct {
	Message string `json:"not implemented yet"`
}

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
