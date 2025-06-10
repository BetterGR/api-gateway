package graph

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/BetterGR/api-gateway/graph/model"
)

// Tool represents an API tool that can be called by an AI model
type Tool struct {
	Name        string                            `json:"name"`        // The name of the tool
	Description string                            `json:"description"` // Description of what the tool does
	Execute     func(map[string]any) (any, error) `json:"-"`           // Function to execute the tool
	Parameters  ToolParameters                    `json:"parameters"`  // Parameters the tool accepts
}

// getContextFromParams extracts the context from params, or creates a new one if not present
func getContextFromParams(params map[string]any) context.Context {
	if ctxParam, ok := params["_context"].(context.Context); ok {
		return ctxParam
	}
	return context.Background()
}

// ToolParameter represents a parameter for a tool
type ToolParameter struct {
	Name        string `json:"name"`        // Name of the parameter
	Type        string `json:"type"`        // Type of the parameter (string, boolean, number, etc.)
	Description string `json:"description"` // Description of the parameter
	Required    bool   `json:"required"`    // Whether the parameter is required
}

// ToolParameters defines the structure for tool parameters
type ToolParameters struct {
	Type       string                   `json:"type"`       // Always "object"
	Properties map[string]ToolParameter `json:"properties"` // Map of parameter names to parameter definitions
	Required   []string                 `json:"required"`   // List of required parameter names
}

// ToolRegistry stores all registered tools
type ToolRegistry struct {
	resolver *Resolver
	tools    map[string]*Tool
}

// NewToolRegistry creates a new tool registry with the given resolver
func NewToolRegistry(resolver *Resolver) *ToolRegistry {
	registry := &ToolRegistry{
		resolver: resolver,
		tools:    make(map[string]*Tool),
	}

	// Register all tools
	registry.registerTools()

	return registry
}

// registerTools registers all API tools with the registry
func (r *ToolRegistry) registerTools() {
	// Student tools
	r.registerStudentTools()

	// Staff tools
	r.registerStaffTools()

	// Course tools
	r.registerCourseTools()

	// Grade tools
	r.registerGradeTools()

	// Homework tools
	r.registerHomeworkTools()

	// Announcement tools
	r.registerAnnouncementTools()

	// ChatBot tool
	r.registerChatBotTool()
}

// registerStudentTools registers all student-related tools
func (r *ToolRegistry) registerStudentTools() {
	// Get Student tool
	r.tools["get_student"] = &Tool{
		Name:        "get_student",
		Description: "Get detailed information about a student by ID",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"id": {
					Name:        "id",
					Type:        "string",
					Description: "The ID of the student",
					Required:    true,
				},
			},
			Required: []string{"id"},
		}, Execute: func(params map[string]any) (any, error) {
			id, ok := params["id"].(string)
			if !ok {
				return nil, fmt.Errorf("id must be a string")
			}

			// Get context with authentication information
			ctx := getContextFromParams(params)

			return r.resolver.Query().Student(ctx, id)
		},
	}

	// Create Student tool
	r.tools["create_student"] = &Tool{
		Name:        "create_student",
		Description: "Create a new student",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"firstName": {
					Name:        "firstName",
					Type:        "string",
					Description: "Student's first name",
					Required:    true,
				},
				"lastName": {
					Name:        "lastName",
					Type:        "string",
					Description: "Student's last name",
					Required:    true,
				},
				"email": {
					Name:        "email",
					Type:        "string",
					Description: "Student's email address",
					Required:    true,
				},
				"phoneNumber": {
					Name:        "phoneNumber",
					Type:        "string",
					Description: "Student's phone number",
					Required:    true,
				},
			},
			Required: []string{"firstName", "lastName", "email", "phoneNumber"},
		}, Execute: func(params map[string]any) (any, error) {
			input := model.NewStudent{
				FirstName:   params["firstName"].(string),
				LastName:    params["lastName"].(string),
				Email:       params["email"].(string),
				PhoneNumber: params["phoneNumber"].(string),
			}

			// Get context with authentication information
			ctx := getContextFromParams(params)

			return r.resolver.Mutation().CreateStudent(ctx, input)
		},
	}
}

// registerStaffTools registers all staff-related tools
func (r *ToolRegistry) registerStaffTools() {
	// Get Staff tool
	r.tools["get_staff"] = &Tool{
		Name:        "get_staff",
		Description: "Get detailed information about a staff member by ID",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"id": {
					Name:        "id",
					Type:        "string",
					Description: "The ID of the staff member",
					Required:    true,
				},
			},
			Required: []string{"id"},
		}, Execute: func(params map[string]any) (any, error) {
			id, ok := params["id"].(string)
			if !ok {
				return nil, fmt.Errorf("id must be a string")
			}

			// Get context with authentication information
			ctx := getContextFromParams(params)

			return r.resolver.Query().Staff(ctx, id)
		},
	}

	// Create Staff tool
	r.tools["create_staff"] = &Tool{
		Name:        "create_staff",
		Description: "Create a new staff member",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"firstName": {
					Name:        "firstName",
					Type:        "string",
					Description: "Staff's first name",
					Required:    true,
				},
				"lastName": {
					Name:        "lastName",
					Type:        "string",
					Description: "Staff's last name",
					Required:    true,
				},
				"email": {
					Name:        "email",
					Type:        "string",
					Description: "Staff's email address",
					Required:    true,
				},
				"phoneNumber": {
					Name:        "phoneNumber",
					Type:        "string",
					Description: "Staff's phone number",
					Required:    true,
				},
				"title": {
					Name:        "title",
					Type:        "string",
					Description: "Staff's title/position",
					Required:    false,
				},
				"office": {
					Name:        "office",
					Type:        "string",
					Description: "Staff's office location",
					Required:    false,
				},
			},
			Required: []string{"firstName", "lastName", "email", "phoneNumber"},
		},
		Execute: func(params map[string]any) (any, error) {
			var title, office *string
			if t, ok := params["title"].(string); ok {
				title = &t
			}
			if o, ok := params["office"].(string); ok {
				office = &o
			}

			input := model.NewStaff{
				FirstName:   params["firstName"].(string),
				LastName:    params["lastName"].(string),
				Email:       params["email"].(string),
				PhoneNumber: params["phoneNumber"].(string),
				Title:       title,
				Office:      office,
			}
			// Get context with authentication information
			ctx := getContextFromParams(params)
			return r.resolver.Mutation().CreateStaff(ctx, input)
		},
	}
}

// registerCourseTools registers all course-related tools
func (r *ToolRegistry) registerCourseTools() {
	// Get Course tool
	r.tools["get_course"] = &Tool{
		Name:        "get_course",
		Description: "Get detailed information about a course by ID",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"id": {
					Name:        "id",
					Type:        "string",
					Description: "The ID of the course",
					Required:    true,
				},
			},
			Required: []string{"id"},
		},
		Execute: func(params map[string]any) (any, error) {
			id, ok := params["id"].(string)
			if !ok {
				return nil, fmt.Errorf("id must be a string")
			}
			// Get context with authentication information
			ctx := getContextFromParams(params)
			return r.resolver.Query().Course(ctx, id)
		},
	}

	// Create Course tool
	r.tools["create_course"] = &Tool{
		Name:        "create_course",
		Description: "Create a new course",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"name": {
					Name:        "name",
					Type:        "string",
					Description: "Course name",
					Required:    true,
				},
				"semester": {
					Name:        "semester",
					Type:        "string",
					Description: "Semester (e.g., 'Fall 2025')",
					Required:    true,
				},
				"description": {
					Name:        "description",
					Type:        "string",
					Description: "Course description",
					Required:    false,
				},
			},
			Required: []string{"name", "semester"},
		},
		Execute: func(params map[string]any) (any, error) {
			var description *string
			if desc, ok := params["description"].(string); ok {
				description = &desc
			}

			input := model.NewCourse{
				Name:        params["name"].(string),
				Semester:    params["semester"].(string),
				Description: description,
			}
			// Get context with authentication information
			ctx := getContextFromParams(params)
			return r.resolver.Mutation().CreateCourse(ctx, input)
		},
	}

	// Get Course Students tool
	r.tools["get_course_students"] = &Tool{
		Name:        "get_course_students",
		Description: "Get all students enrolled in a specific course",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"courseId": {
					Name:        "courseId",
					Type:        "string",
					Description: "The ID of the course",
					Required:    true,
				},
			},
			Required: []string{"courseId"},
		},
		Execute: func(params map[string]any) (any, error) {
			courseId, ok := params["courseId"].(string)
			if !ok {
				return nil, fmt.Errorf("courseId must be a string")
			}
			// Get context with authentication information
			ctx := getContextFromParams(params)
			return r.resolver.Query().CourseStudents(ctx, courseId)
		},
	}
}

// registerGradeTools registers all grade-related tools
func (r *ToolRegistry) registerGradeTools() {
	// Get Student Course Grades tool
	r.tools["get_student_course_grades"] = &Tool{
		Name:        "get_student_course_grades",
		Description: "Get all grades for a specific student in a specific course and semester",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"studentId": {
					Name:        "studentId",
					Type:        "string",
					Description: "The ID of the student",
					Required:    true,
				},
				"courseId": {
					Name:        "courseId",
					Type:        "string",
					Description: "The ID of the course",
					Required:    true,
				},
				"semester": {
					Name:        "semester",
					Type:        "string",
					Description: "The semester (e.g., 'Fall 2025')",
					Required:    true,
				},
			},
			Required: []string{"studentId", "courseId", "semester"},
		},
		Execute: func(params map[string]any) (any, error) {
			studentId := params["studentId"].(string)
			courseId := params["courseId"].(string)
			semester := params["semester"].(string)
			// Get context with authentication information
			ctx := getContextFromParams(params)
			return r.resolver.Query().StudentCourseGrades(ctx, studentId, courseId, semester)
		},
	}

	// Create Grade tool
	r.tools["create_grade"] = &Tool{
		Name:        "create_grade",
		Description: "Create a new grade entry for a student",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"studentId": {
					Name:        "studentId",
					Type:        "string",
					Description: "The ID of the student",
					Required:    true,
				},
				"courseId": {
					Name:        "courseId",
					Type:        "string",
					Description: "The ID of the course",
					Required:    true,
				},
				"semester": {
					Name:        "semester",
					Type:        "string",
					Description: "The semester (e.g., 'Fall 2025')",
					Required:    true,
				},
				"gradeType": {
					Name:        "gradeType",
					Type:        "string",
					Description: "Type of grade (e.g., 'quiz', 'exam', 'homework')",
					Required:    true,
				},
				"itemId": {
					Name:        "itemId",
					Type:        "string",
					Description: "ID of the graded item (e.g., 'Quiz 1', 'Midterm')",
					Required:    true,
				},
				"gradeValue": {
					Name:        "gradeValue",
					Type:        "string",
					Description: "The actual grade value (e.g., '95', 'A-')",
					Required:    true,
				},
				"gradedBy": {
					Name:        "gradedBy",
					Type:        "string",
					Description: "ID of the staff member who graded the item",
					Required:    false,
				},
				"comments": {
					Name:        "comments",
					Type:        "string",
					Description: "Comments on the grade",
					Required:    false,
				},
			},
			Required: []string{"studentId", "courseId", "semester", "gradeType", "itemId", "gradeValue"},
		},
		Execute: func(params map[string]any) (any, error) {
			var gradedBy *string
			var comments *string

			if g, ok := params["gradedBy"].(string); ok {
				gradedBy = &g
			}

			if c, ok := params["comments"].(string); ok {
				comments = &c
			}

			input := model.NewGrade{
				StudentID:  params["studentId"].(string),
				CourseID:   params["courseId"].(string),
				Semester:   params["semester"].(string),
				GradeType:  params["gradeType"].(string),
				ItemID:     params["itemId"].(string),
				GradeValue: params["gradeValue"].(string),
				GradedBy:   gradedBy,
				Comments:   comments,
			}
			// Get context with authentication information
			ctx := getContextFromParams(params)
			return r.resolver.Mutation().CreateGrade(ctx, input)
		},
	}
}

// registerHomeworkTools registers all homework-related tools
func (r *ToolRegistry) registerHomeworkTools() {
	// Get Course Homework tool
	r.tools["get_course_homework"] = &Tool{
		Name:        "get_course_homework",
		Description: "Get all homework assignments for a specific course",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"courseId": {
					Name:        "courseId",
					Type:        "string",
					Description: "The ID of the course",
					Required:    true,
				},
			},
			Required: []string{"courseId"},
		},
		Execute: func(params map[string]any) (any, error) {
			courseId, ok := params["courseId"].(string)
			if !ok {
				return nil, fmt.Errorf("courseId must be a string")
			}
			// Get context with authentication information
			ctx := getContextFromParams(params)
			return r.resolver.Query().HomeworkByCourse(ctx, courseId)
		},
	}

	// Create Homework tool
	r.tools["create_homework"] = &Tool{
		Name:        "create_homework",
		Description: "Create a new homework assignment for a course",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"courseId": {
					Name:        "courseId",
					Type:        "string",
					Description: "The ID of the course",
					Required:    true,
				},
				"title": {
					Name:        "title",
					Type:        "string",
					Description: "Title of the homework assignment",
					Required:    true,
				},
				"description": {
					Name:        "description",
					Type:        "string",
					Description: "Detailed description of the homework",
					Required:    true,
				},
				"workflow": {
					Name:        "workflow",
					Type:        "string",
					Description: "Workflow/instructions for completing the homework",
					Required:    true,
				},
				"dueDate": {
					Name:        "dueDate",
					Type:        "string",
					Description: "Due date in ISO format (e.g., '2025-07-15T23:59:59Z')",
					Required:    true,
				},
			},
			Required: []string{"courseId", "title", "description", "workflow", "dueDate"},
		},
		Execute: func(params map[string]any) (any, error) {
			input := model.NewHomework{
				CourseID:    params["courseId"].(string),
				Title:       params["title"].(string),
				Description: params["description"].(string),
				Workflow:    params["workflow"].(string),
				DueDate:     params["dueDate"].(string),
			}
			// Get context with authentication information
			ctx := getContextFromParams(params)
			return r.resolver.Mutation().CreateHomework(ctx, input)
		},
	}
}

// registerAnnouncementTools registers all announcement-related tools
func (r *ToolRegistry) registerAnnouncementTools() {
	// Get Course Announcements tool
	r.tools["get_course_announcements"] = &Tool{
		Name:        "get_course_announcements",
		Description: "Get all announcements for a specific course",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"courseId": {
					Name:        "courseId",
					Type:        "string",
					Description: "The ID of the course",
					Required:    true,
				},
			},
			Required: []string{"courseId"},
		},
		Execute: func(params map[string]any) (any, error) {
			courseId, ok := params["courseId"].(string)
			if !ok {
				return nil, fmt.Errorf("courseId must be a string")
			}
			// Get context with authentication information
			ctx := getContextFromParams(params)
			return r.resolver.Query().AnnouncementsByCourse(ctx, courseId)
		},
	}

	// Create Announcement tool
	r.tools["create_announcement"] = &Tool{
		Name:        "create_announcement",
		Description: "Create a new announcement for a course",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"courseId": {
					Name:        "courseId",
					Type:        "string",
					Description: "The ID of the course",
					Required:    true,
				},
				"title": {
					Name:        "title",
					Type:        "string",
					Description: "Title of the announcement",
					Required:    true,
				},
				"content": {
					Name:        "content",
					Type:        "string",
					Description: "Content of the announcement",
					Required:    true,
				},
			},
			Required: []string{"courseId", "title", "content"},
		},
		Execute: func(params map[string]any) (any, error) {
			input := model.NewAnnouncement{
				CourseID: params["courseId"].(string),
				Title:    params["title"].(string),
				Content:  params["content"].(string),
			}
			// Get context with authentication information
			ctx := getContextFromParams(params)
			return r.resolver.Mutation().CreateAnnouncement(ctx, input)
		},
	}
}

// registerChatBotTool registers the ChatBot tool
func (r *ToolRegistry) registerChatBotTool() {
	// Process Chat Message tool
	r.tools["process_chat_message"] = &Tool{
		Name:        "process_chat_message",
		Description: "Process a chat message with context and history",
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolParameter{
				"newMessage": {
					Name:        "newMessage",
					Type:        "string",
					Description: "The new message to process",
					Required:    true,
				},
				"chatHistory": {
					Name:        "chatHistory",
					Type:        "array",
					Description: "Array of previous chat messages in the format [{role: 'user|assistant', content: 'message'}]",
					Required:    true,
				},
				"userId": {
					Name:        "userId",
					Type:        "string",
					Description: "ID of the user",
					Required:    false,
				},
				"userRole": {
					Name:        "userRole",
					Type:        "string",
					Description: "Role of the user (e.g., 'student', 'staff')",
					Required:    false,
				},
				"courseId": {
					Name:        "courseId",
					Type:        "string",
					Description: "ID of the course if the chat is course-specific",
					Required:    false,
				},
				"sessionId": {
					Name:        "sessionId",
					Type:        "string",
					Description: "ID of the chat session",
					Required:    false,
				},
			},
			Required: []string{"newMessage", "chatHistory"},
		},
		Execute: func(params map[string]any) (any, error) {
			// Convert chat history to the expected format
			chatHistory, ok := params["chatHistory"].([]interface{})
			if !ok {
				return nil, fmt.Errorf("chatHistory must be an array")
			}

			var modelChatHistory []*model.ChatMessageInput
			for _, msg := range chatHistory {
				msgMap, ok := msg.(map[string]interface{})
				if !ok {
					continue
				}

				role, _ := msgMap["role"].(string)
				content, _ := msgMap["content"].(string)

				if role != "" && content != "" {
					modelChatHistory = append(modelChatHistory, &model.ChatMessageInput{
						Role:    role,
						Content: content,
					})
				}
			}
			// Create context if provided
			var chatContext *model.ChatContextInput
			if params["userId"] != nil || params["userRole"] != nil || params["courseId"] != nil || params["sessionId"] != nil {
				chatContext = &model.ChatContextInput{}

				if userId, ok := params["userId"].(string); ok {
					chatContext.UserID = userId
				}

				if userRole, ok := params["userRole"].(string); ok {
					chatContext.UserRole = userRole
				}

				if courseId, ok := params["courseId"].(string); ok {
					chatContext.CourseID = courseId
				}

				if sessionId, ok := params["sessionId"].(string); ok {
					chatContext.SessionID = sessionId
				}
			}

			input := model.ChatHistoryInput{
				NewMessage:  params["newMessage"].(string),
				ChatHistory: modelChatHistory,
				Context:     chatContext,
			}
			// Get context with authentication information
			ctx := getContextFromParams(params)
			// Use the resolver's mutation handler to process the chat message
			mutationResolver := r.resolver
			return mutationResolver.ChatBotClient.ProcessMessage(ctx, &input)
		},
	}
}

// GetTool returns a tool by name
func (r *ToolRegistry) GetTool(name string) (*Tool, bool) {
	tool, exists := r.tools[name]
	return tool, exists
}

// GetAllTools returns a list of all available tools
func (r *ToolRegistry) GetAllTools() []*Tool {
	tools := make([]*Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		tools = append(tools, tool)
	}
	return tools
}

// ExecuteTool executes a tool by name with the given parameters and authentication token
func (r *ToolRegistry) ExecuteTool(name string, params map[string]any, authToken string) (any, error) {
	tool, exists := r.tools[name]
	if !exists {
		return nil, fmt.Errorf("tool %s not found", name)
	}

	// Create a context with the authentication token
	ctx := context.Background()
	if authToken != "" {
		// Store the token in the context
		ctx = context.WithValue(ctx, AuthTokenKey, authToken)
		// Create an auth context with the token for passing to microservices
		ctx = r.resolver.CreateAuthContext(ctx)
	}

	// Store the context in the params for Execute functions to use
	params["_context"] = ctx

	// Execute the tool with the provided parameters
	return tool.Execute(params)
}

// GetToolsAsJSON returns all tools in JSON Schema format for LLM consumption
func (r *ToolRegistry) GetToolsAsJSON() (string, error) {
	tools := r.GetAllTools()

	// Convert tools to a format suitable for LLMs
	type LLMTool struct {
		Name        string         `json:"name"`
		Description string         `json:"description"`
		Parameters  ToolParameters `json:"parameters"`
	}

	llmTools := make([]LLMTool, len(tools))
	for i, tool := range tools {
		llmTools[i] = LLMTool{
			Name:        tool.Name,
			Description: tool.Description,
			Parameters:  tool.Parameters,
		}
	}

	// Marshal to JSON
	jsonBytes, err := json.Marshal(llmTools)
	if err != nil {
		return "", fmt.Errorf("error marshaling tools to JSON: %w", err)
	}

	return string(jsonBytes), nil
}
