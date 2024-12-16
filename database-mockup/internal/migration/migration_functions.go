package migration

import (
	"bookmark-database/table"
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
	s := v.(string)
	return &s
}

func nilableUint64Ptr(v interface{}) *uint64 {
	if v == nil {
		return nil
	}
	i := uint64(v.(float64))
	return &i
}

// Migration functions
func MigrateUsers(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		user := &table.User{
			Id:        toUint64Ptr(d["id"].(float64)),
			Oid:       toStringPtr(d["oid"].(string)),
			Firstname: toStringPtr(d["firstname"].(string)),
			Lastname:  toStringPtr(d["lastname"].(string)),
			Email:     toStringPtr(d["email"].(string)),
			PhotoUrl:  nilableStringPtr(d["photo_url"]),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}
	}
	return nil
}

func MigrateFields(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		field := &table.Field{
			Id:        toUint64Ptr(d["id"].(float64)),
			Name:      toStringPtr(d["name"].(string)),
			ImageUrl:  nilableStringPtr(d["image_url"]),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(field).Error; err != nil {
			return err
		}
	}
	return nil
}

func MigrateModules(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		module := &table.Module{
			Id:          toUint64Ptr(d["id"].(float64)),
			Title:       toStringPtr(d["title"].(string)),
			Description: nilableStringPtr(d["description"]),
			ImageUrl:    nilableStringPtr(d["image_url"]),
			CreatedAt:   toTimePtr(d["created_at"].(string)),
			UpdatedAt:   toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(module).Error; err != nil {
			return err
		}
	}
	return nil
}

func MigrateCourses(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		course := &table.Course{
			Id:        toUint64Ptr(d["id"].(float64)),
			Name:      toStringPtr(d["name"].(string)),
			FieldId:   toUint64Ptr(d["field_id"].(float64)),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(course).Error; err != nil {
			return err
		}
	}
	return nil
}

func MigrateSteps(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		step := &table.Step{
			Id:          toUint64Ptr(d["id"].(float64)),
			ModuleId:    toUint64Ptr(d["module_id"].(float64)),
			Title:       toStringPtr(d["title"].(string)),
			Description: nilableStringPtr(d["description"]),
			Gems:        toInt64Ptr(d["gems"].(float64)),
			Content:     nilableStringPtr(d["content"]),
			Outcome:     nilableStringPtr(d["outcome"]),
			Check:       toBoolPtr(d["check"].(bool)),
			CreatedAt:   toTimePtr(d["created_at"].(string)),
			UpdatedAt:   toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(step).Error; err != nil {
			return err
		}
	}
	return nil
}

func MigrateStepAuthors(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		author := &table.StepAuthor{
			StepId: toUint64Ptr(d["step_id"].(float64)),
			UserId: toUint64Ptr(d["user_id"].(float64)),
		}
		if err := tx.Create(author).Error; err != nil {
			return err
		}
	}
	return nil
}

func MigrateStepEvaluates(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		evaluate := &table.StepEvaluate{
			Id:        toUint64Ptr(d["id"].(float64)),
			StepId:    toUint64Ptr(d["step_id"].(float64)),
			Order:     toIntPtr(d["order"].(float64)),
			Type:      toStringPtr(d["type"].(string)),
			Prompt:    toStringPtr(d["prompt"].(string)),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(evaluate).Error; err != nil {
			return err
		}
	}
	return nil
}

func MigrateUserEvaluates(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		evaluate := &table.UserEvaluate{
			UserId:         toUint64Ptr(d["user_id"].(float64)),
			StepEvaluateId: toUint64Ptr(d["step_evaluate_id"].(float64)),
			Content:        toStringPtr(d["content"].(string)),
			Check:          toBoolPtr(d["check"].(bool)),
			Comment:        nilableStringPtr(d["comment"]),
			CreatedAt:      toTimePtr(d["created_at"].(string)),
			UpdatedAt:      toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(evaluate).Error; err != nil {
			return err
		}
	}
	return nil
}

func MigrateStepComments(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		comment := &table.StepComment{
			Id:        toUint64Ptr(d["id"].(float64)),
			StepId:    toUint64Ptr(d["step_id"].(float64)),
			UserId:    toUint64Ptr(d["user_id"].(float64)),
			Content:   toStringPtr(d["content"].(string)),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(comment).Error; err != nil {
			return err
		}
	}
	return nil
}

func MigrateStepCommentUpvotes(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		upvote := &table.StepCommentUpvote{
			StepCommentId: toUint64Ptr(d["step_comment_id"].(float64)),
			UserId:        toUint64Ptr(d["user_id"].(float64)),
			CreatedAt:     toTimePtr(d["created_at"].(string)),
			UpdatedAt:     toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(upvote).Error; err != nil {
			return err
		}
	}
	return nil
}

func MigrateUserPasses(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		pass := &table.UserPass{
			Id:        toUint64Ptr(d["id"].(float64)),
			UserId:    toUint64Ptr(d["user_id"].(float64)),
			Type:      toStringPtr(d["type"].(string)),
			StepId:    nilableUint64Ptr(d["step_id"]),
			CourseId:  nilableUint64Ptr(d["course_id"]),
			ModuleId:  nilableUint64Ptr(d["module_id"]),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(pass).Error; err != nil {
			return err
		}
	}
	return nil
}

func MigrateEnrols(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		enrol := &table.Enroll{
			Id:        toUint64Ptr(d["id"].(float64)),
			UserId:    toUint64Ptr(d["user_id"].(float64)),
			CourseId:  toUint64Ptr(d["course_id"].(float64)),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(enrol).Error; err != nil {
			return err
		}
	}
	return nil
}

func MigrateCourseContents(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		content := &table.CourseContent{
			CourseId:  toUint64Ptr(d["course_id"].(float64)),
			Order:     toInt64Ptr(d["order"].(float64)),
			Type:      toStringPtr(d["type"].(string)),
			Text:      nilableStringPtr(d["text"]),
			ModuleId:  nilableUint64Ptr(d["module_id"]),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(content).Error; err != nil {
			return err
		}
	}
	return nil
}

func MigrateArticles(tx *gorm.DB, data []map[string]interface{}) error {
	for _, d := range data {
		article := &table.Article{
			Id:        toUint64Ptr(d["id"].(float64)),
			Title:     toStringPtr(d["title"].(string)),
			Href:      toStringPtr(d["href"].(string)),
			CreatedAt: toTimePtr(d["created_at"].(string)),
			UpdatedAt: toTimePtr(d["updated_at"].(string)),
		}
		if err := tx.Create(article).Error; err != nil {
			return err
		}
	}
	return nil
}
