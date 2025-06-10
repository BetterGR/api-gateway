package graph

import (
	"context"
	"fmt"
	"os"

	coursespb "github.com/BetterGR/courses-microservice/protos"
	gradespb "github.com/BetterGR/grades-microservice/protos"
	staffpb "github.com/BetterGR/staff-microservice/protos"
	studentspb "github.com/BetterGR/students-microservice/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	StudentsClient studentspb.StudentsServiceClient
	StaffClient    staffpb.StaffServiceClient
	CoursesClient  coursespb.CoursesServiceClient
	GradesClient   gradespb.GradesServiceClient
	ChatBotClient  *ChatBotClient
	ToolRegistry   *ToolRegistry

	// Store connection objects to properly close them
	studentsConn *grpc.ClientConn
	staffConn    *grpc.ClientConn
	coursesConn  *grpc.ClientConn
	gradesConn   *grpc.ClientConn
}

// Close properly closes all gRPC connections
func (r *Resolver) Close() {
	if r.studentsConn != nil {
		r.studentsConn.Close()
	}
	if r.staffConn != nil {
		r.staffConn.Close()
	}
	if r.coursesConn != nil {
		r.coursesConn.Close()
	}
	if r.gradesConn != nil {
		r.gradesConn.Close()
	}
}

// NewResolver creates a new resolver with all the necessary gRPC clients
func NewResolver() (*Resolver, error) {
	// Get microservice endpoints from environment variables or use defaults
	gradesEndpoint := getEnvOrDefault("GRADES_PORT", "localhost:50051")
	studentsEndpoint := getEnvOrDefault("STUDENTS_PORT", "localhost:50052")
	// homeworkEndpoint := getEnvOrDefault("HOMEWORK_PORT", "localhost:50053")
	coursesEndpoint := getEnvOrDefault("COURSES_PORT", "localhost:50054")
	staffEndpoint := getEnvOrDefault("STAFF_PORT", "localhost:50055")

	// Setup connection to Students microservice
	studentsConn, err := grpc.NewClient(
		studentsEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to students service: %w", err)
	}
	studentsClient := studentspb.NewStudentsServiceClient(studentsConn)

	// Setup connection to Staff microservice
	staffConn, err := grpc.NewClient(
		staffEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to staff service: %w", err)
	}
	staffClient := staffpb.NewStaffServiceClient(staffConn)

	// Setup connection to Courses microservice
	coursesConn, err := grpc.NewClient(
		coursesEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to courses service: %w", err)
	}
	coursesClient := coursespb.NewCoursesServiceClient(coursesConn)

	// Setup connection to Grades microservice
	gradesConn, err := grpc.NewClient(
		gradesEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to grades service: %w", err)
	}
	gradesClient := gradespb.NewGradesServiceClient(gradesConn)

	// Note: Homework service is not used directly as it's probably part of the courses service

	// Initialize the ChatBot client
	chatBotClient := NewChatBotClient()

	resolver := &Resolver{
		StudentsClient: studentsClient,
		StaffClient:    staffClient,
		CoursesClient:  coursesClient,
		GradesClient:   gradesClient,
		ChatBotClient:  chatBotClient,
		studentsConn:   studentsConn,
		staffConn:      staffConn,
		coursesConn:    coursesConn,
		gradesConn:     gradesConn,
	}

	// Initialize the tool registry with this resolver
	resolver.ToolRegistry = NewToolRegistry(resolver)

	return resolver, nil
}

// Helper function to get environment variable with fallback
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// CreateAuthContext creates a new context with authentication metadata from the GraphQL context
func (r *Resolver) CreateAuthContext(ctx context.Context) context.Context {
	token := GetAuthToken(ctx)
	if token == "" {
		return ctx
	}

	// Create gRPC metadata with the authorization token
	md := metadata.Pairs("authorization", "Bearer "+token)
	return metadata.NewOutgoingContext(ctx, md)
}

// GetAuthTokenForRequest extracts the auth token for use in request messages
func (r *Resolver) GetAuthTokenForRequest(ctx context.Context) string {
	return GetAuthToken(ctx)
}
