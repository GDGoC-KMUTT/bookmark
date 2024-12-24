package migration

import (
	"backend/internals/db/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Helper functions
func toUint64Ptr(v float64) *uint64 {
	i := uint64(v)
	return &i
}

func toInt64Ptr(v float64) *int64 {
	i := int64(v)
	return &i
}

func toIntPtr(v float64) *int {
	i := int(v)
	return &i
}

func toStringPtr(v string) *string {
	return &v
}

func toBoolPtr(v bool) *bool {
	return &v
}

func toTimePtr(v string) *time.Time {
	t, _ := time.Parse(time.RFC3339, v)
	return &t
}

func nilableStringPtr(v interface{}) *string {
	if v == nil {
		return nil
	}
	switch t := v.(type) {
	case string:
		return &t
	case bool:
		s := fmt.Sprintf("%v", t)
		return &s
	default:
		return nil
	}
}

func nilableUint64Ptr(v interface{}) *uint64 {
	if v == nil {
		return nil
	}
	if f, ok := v.(float64); ok {
		i := uint64(f)
		return &i
	}
	return nil
}

func safeCastToString(v interface{}) (string, error) {
	if v == nil {
		return "", fmt.Errorf("value is nil")
	}
	if s, ok := v.(string); ok {
		return s, nil
	}
	return "", fmt.Errorf("value is not a string")
}

func safeCastToFloat64(v interface{}) (float64, error) {
	if v == nil {
		return 0, fmt.Errorf("value is nil")
	}
	if f, ok := v.(float64); ok {
		return f, nil
	}
	return 0, fmt.Errorf("value is not a float64")
}

func safeCastToBool(v interface{}) (bool, error) {
	if v == nil {
		return false, fmt.Errorf("value is nil")
	}
	if b, ok := v.(bool); ok {
		return b, nil
	}
	return false, fmt.Errorf("value is not a bool")
}

// Migration functions
func MigrateUsers(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		id, err := safeCastToFloat64(d["id"])
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}

		user := &models.User{
			Id:        toUint64Ptr(id),
			Oid:       toStringPtr(d["oid"].(string)),
			Firstname: toStringPtr(d["firstname"].(string)),
			Lastname:  toStringPtr(d["lastname"].(string)),
			Email:     toStringPtr(d["email"].(string)),
			PhotoUrl:  nilableStringPtr(d["photo_url"]),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(user).Error; err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
	}
	return nil
}

func MigrateFieldTypes(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		id, err := safeCastToFloat64(d["id"])
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}

		field := &models.FieldType{
			Id:        toUint64Ptr(id),
			Name:      toStringPtr(d["name"].(string)),
			ImageUrl:  nilableStringPtr(d["image_url"]),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(field).Error; err != nil {
			return fmt.Errorf("failed to create field: %w", err)
		}
	}
	return nil
}

func MigrateModules(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		id, err := safeCastToFloat64(d["id"])
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}

		module := &models.Module{
			Id:          toUint64Ptr(id),
			Title:       toStringPtr(d["title"].(string)),
			Description: nilableStringPtr(d["description"]),
			ImageUrl:    nilableStringPtr(d["image_url"]),
			CreatedAt:   toTimePtr(d["created_at"].(string)),
			UpdatedAt:   toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(module).Error; err != nil {
			return fmt.Errorf("failed to create module: %w", err)
		}
	}
	return nil
}

func MigrateSteps(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		id, err := safeCastToFloat64(d["id"])
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}

		moduleId, err := safeCastToFloat64(d["module_id"])
		if err != nil {
			return fmt.Errorf("invalid module_id: %w", err)
		}
		
		step := &models.Step{
			Id:          toUint64Ptr(id),
			ModuleId:    toUint64Ptr(moduleId),
			Title:       toStringPtr(d["title"].(string)),
			Description: nilableStringPtr(d["description"]),
			Content:     nilableStringPtr(d["content"]),
			Outcome:     nilableStringPtr(d["outcome"]),
			Check:       nilableStringPtr(d["check"]),
			Error:       nilableStringPtr(d["error"]),
			CreatedAt:   toTimePtr(d["created_at"].(string)),
			UpdatedAt:   toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(step).Error; err != nil {
			return fmt.Errorf("failed to create step: %w", err)
		}
	}
	return nil
}

func MigrateStepAuthors(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		stepId, err := safeCastToFloat64(d["step_id"])
		if err != nil {
			return fmt.Errorf("invalid step_id: %w", err)
		}

		userId, err := safeCastToFloat64(d["user_id"])
		if err != nil {
			return fmt.Errorf("invalid user_id: %w", err)
		}

		author := &models.StepAuthor{
			StepId: toUint64Ptr(stepId),
			UserId: toUint64Ptr(userId),
		}
		if err := tx.Create(author).Error; err != nil {
			return fmt.Errorf("failed to create step author: %w", err)
		}
	}
	return nil
}

func MigrateStepEvaluates(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		evaluate := &models.StepEvaluate{
			Id:          toUint64Ptr(d["id"].(float64)),
			StepId:      toUint64Ptr(d["step_id"].(float64)),
			Gem:         toIntPtr(d["gem"].(float64)),
			Order:       toIntPtr(d["order"].(float64)),
			Type:        toStringPtr(d["type"].(string)),
			Question:    toStringPtr(d["question"].(string)),
			Instruction: toStringPtr(d["instruction"].(string)),
			CreatedAt:   toTimePtr(d["created_at"].(string)),
			UpdatedAt:   toTimePtr(d["updated_at"].(string)),
		}

		if err := tx.Create(evaluate).Error; err != nil {
			return fmt.Errorf("failed to create step evaluate: %w", err)
		}
	}
	return nil
}

func MigrateUserEvaluates(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {

		var pass bool
		if passVal, exists := d["pass"]; exists {
			pass = passVal.(bool)
		}

		evaluate := &models.UserEvaluate{
			UserId:         toUint64Ptr(d["user_id"].(float64)),
			StepEvaluateId: toUint64Ptr(d["step_evaluate_id"].(float64)),
			Content:        toStringPtr(d["content"].(string)),
			Pass:           toBoolPtr(pass),
			Comment:        nilableStringPtr(d["comment"]),
			CreatedAt:      toTimePtr(d["created_at"].(string)),
			UpdatedAt:      toTimePtr(d["updated_at"].(string)),
		}

		if err := tx.Create(evaluate).Error; err != nil {
			return fmt.Errorf("failed to create user evaluate: %w", err)
		}
	}
	return nil
}

func MigrateStepComments(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		id, err := safeCastToFloat64(d["id"])
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}

		stepId, err := safeCastToFloat64(d["step_id"])
		if err != nil {
			return fmt.Errorf("invalid step_id: %w", err)
		}

		userId, err := safeCastToFloat64(d["user_id"])
		if err != nil {
			return fmt.Errorf("invalid user_id: %w", err)
		}

		comment := &models.StepComment{
			Id:        toUint64Ptr(id),
			StepId:    toUint64Ptr(stepId),
			UserId:    toUint64Ptr(userId),
			Content:   toStringPtr(d["content"].(string)),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(comment).Error; err != nil {
			return fmt.Errorf("failed to create step comment: %w", err)
		}
	}
	return nil
}

func MigrateStepCommentUpvotes(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		stepCommentId, err := safeCastToFloat64(d["step_comment_id"])
		if err != nil {
			return fmt.Errorf("invalid step_comment_id: %w", err)
		}

		userId, err := safeCastToFloat64(d["user_id"])
		if err != nil {
			return fmt.Errorf("invalid user_id: %w", err)
		}

		upvote := &models.StepCommentUpvote{
			StepCommentId: toUint64Ptr(stepCommentId),
			UserId:        toUint64Ptr(userId),
			CreatedAt:     toTimePtr(d["created_at"].(string)),
			UpdatedAt:     toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(upvote).Error; err != nil {
			return fmt.Errorf("failed to create step comment upvote: %w", err)
		}
	}
	return nil
}

func MigrateUserPasses(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		id, err := safeCastToFloat64(d["id"])
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}

		userId, err := safeCastToFloat64(d["user_id"])
		if err != nil {
			return fmt.Errorf("invalid user_id: %w", err)
		}

		pass := &models.UserPass{
			Id:        toUint64Ptr(id),
			UserId:    toUint64Ptr(userId),
			Type:      toStringPtr(d["type"].(string)),
			StepId:    nilableUint64Ptr(d["step_id"]),
			CourseId:  nilableUint64Ptr(d["course_id"]),
			ModuleId:  nilableUint64Ptr(d["module_id"]),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(pass).Error; err != nil {
			return fmt.Errorf("failed to create user pass: %w", err)
		}
	}
	return nil
}

func MigrateEnrolls(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		id, err := safeCastToFloat64(d["id"])
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}

		userId, err := safeCastToFloat64(d["user_id"])
		if err != nil {
			return fmt.Errorf("invalid user_id: %w", err)
		}

		courseId, err := safeCastToFloat64(d["course_id"])
		if err != nil {
			return fmt.Errorf("invalid course_id: %w", err)
		}

		enrol := &models.Enroll{
			Id:        toUint64Ptr(id),
			UserId:    toUint64Ptr(userId),
			CourseId:  toUint64Ptr(courseId),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(enrol).Error; err != nil {
			return fmt.Errorf("failed to create enrol: %w", err)
		}
	}
	return nil
}

func MigrateCourseContents(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		courseId, err := safeCastToFloat64(d["course_id"])
		if err != nil {
			return fmt.Errorf("invalid course_id: %w", err)
		}

		order, err := safeCastToFloat64(d["order"])
		if err != nil {
			return fmt.Errorf("invalid order: %w", err)
		}

		content := &models.CourseContent{
			CourseId:  toUint64Ptr(courseId),
			Order:     toInt64Ptr(order),
			Type:      toStringPtr(d["type"].(string)),
			Text:      nilableStringPtr(d["text"]),
			ModuleId:  nilableUint64Ptr(d["module_id"]),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(content).Error; err != nil {
			return fmt.Errorf("failed to create course content: %w", err)
		}
	}
	return nil
}

func MigrateArticles(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		id, err := safeCastToFloat64(d["id"])
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}

		article := &models.Article{
			Id:        toUint64Ptr(id),
			Title:     toStringPtr(d["title"].(string)),
			Href:      toStringPtr(d["href"].(string)),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(article).Error; err != nil {
			return fmt.Errorf("failed to create article: %w", err)
		}
	}
	return nil
}

func MigrateCourses(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		id, err := safeCastToFloat64(d["id"])
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}

		fieldId, err := safeCastToFloat64(d["field_id"])
		if err != nil {
			return fmt.Errorf("invalid field_id: %w", err)
		}

		course := &models.Course{
			Id:        toUint64Ptr(id),
			Name:      toStringPtr(d["name"].(string)),
			FieldId:   toUint64Ptr(fieldId),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(course).Error; err != nil {
			return fmt.Errorf("failed to create course: %w", err)
		}
	}
	return nil
}

func MigrateUserActivity(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		userId, err := safeCastToFloat64(d["user_id"])
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		stepId, err := safeCastToFloat64(d["step_id"])
		if err != nil {
			return fmt.Errorf("invalid step_id: %w", err)
		}

		userActivity := &models.UserActivity{
			UserId:    toUint64Ptr(userId),
			StepId:    toUint64Ptr(stepId),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(userActivity).Error; err != nil {
			return fmt.Errorf("failed to create user activity: %w", err)
		}
	}
	return nil
}
