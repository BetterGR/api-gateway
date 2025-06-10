package model

// ChatHistoryInput represents the complete input for processing a chat message
type ChatHistoryInput struct {
	NewMessage  string              `json:"newMessage"`
	ChatHistory []*ChatMessageInput `json:"chatHistory"`
	Context     *ChatContextInput   `json:"context,omitempty"`
}

// ChatMessageInput represents a single message in the chat history
type ChatMessageInput struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatContextInput provides contextual information about the chat session
type ChatContextInput struct {
	UserID    string `json:"userId,omitempty"`
	UserRole  string `json:"userRole,omitempty"`
	CourseID  string `json:"courseId,omitempty"`
	SessionID string `json:"sessionId,omitempty"`
}
