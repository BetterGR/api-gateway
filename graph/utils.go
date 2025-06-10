package graph

import (
	"fmt"

	"github.com/BetterGR/api-gateway/graph/model"
	gradespb "github.com/BetterGR/grades-microservice/protos"
)

func convertGradesToGraphQL(grades []*gradespb.SingleGrade) []*model.Grade {
	result := make([]*model.Grade, len(grades))

	for i, g := range grades {
		// For now, GradedBy contains the timestamp, so we use it for both timestamp fields
		timestamp := g.GradedBy
		gradedBy := "System" // Default value since we're using the field for timestamp

		// Debug: Print what we're receiving
		fmt.Printf("DEBUG: Received grade with GradedBy (timestamp): %s\n", timestamp)

		result[i] = &model.Grade{
			ID:         g.GradeID,
			StudentID:  g.StudentID,
			CourseID:   g.CourseID,
			Semester:   g.Semester,
			GradeType:  g.GradeType,
			ItemID:     g.ItemID,
			GradeValue: g.GradeValue,
			GradedBy:   &gradedBy,
			Comments:   &g.Comments,
			GradedAt:   timestamp,
			UpdatedAt:  timestamp,
		}
	}

	return result
}
