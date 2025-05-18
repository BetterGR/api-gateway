package graph

import (
	"time"

	"github.com/BetterGR/api-gateway/graph/model"
	gradespb "github.com/BetterGR/grades-microservice/protos"
)

func convertGradesToGraphQL(grades []*gradespb.SingleGrade) []*model.Grade {
	result := make([]*model.Grade, len(grades))
	now := time.Now().Format(time.RFC3339)

	for i, g := range grades {
		result[i] = &model.Grade{
			ID:         g.GradeID,
			StudentID:  g.StudentID,
			CourseID:   g.CourseID,
			Semester:   g.Semester,
			GradeType:  g.GradeType,
			ItemID:     g.ItemID,
			GradeValue: g.GradeValue,
			GradedBy:   &g.GradedBy,
			Comments:   &g.Comments,
			GradedAt:   now, // In a real implementation, this would come from the microservice
			UpdatedAt:  now,
		}
	}

	return result
}
