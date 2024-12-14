package db

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

var UserModel *gorm.DB
var CourseModel *gorm.DB
var ArticleModel *gorm.DB
var CourseContentModel *gorm.DB
var EnrollModel *gorm.DB
var FieldTypeModel *gorm.DB
var ModuleModel *gorm.DB
var StepModel *gorm.DB
var StepAuthorModel *gorm.DB
var StepCommentModel *gorm.DB
var StepCommentUpvoteModel *gorm.DB
var StepEvaluateModel *gorm.DB
var UserActivityModel *gorm.DB
var UserEvaluateModel *gorm.DB
var UserPassModel *gorm.DB

func AssignModel() {
	UserModel = Gorm.Model(new(models.User))
	CourseModel = Gorm.Model(new(models.Course))
	ArticleModel = Gorm.Model(new(models.Article))
	CourseContentModel = Gorm.Model(new(models.CourseContent))
	EnrollModel = Gorm.Model(new(models.Enroll))
	FieldTypeModel = Gorm.Model(new(models.FieldType))
	ModuleModel = Gorm.Model(new(models.Module))
	StepModel = Gorm.Model(new(models.Step))
	StepAuthorModel = Gorm.Model(new(models.StepAuthor))
	StepCommentModel = Gorm.Model(new(models.StepComment))
	StepCommentUpvoteModel = Gorm.Model(new(models.StepCommentUpvote))
	StepEvaluateModel = Gorm.Model(new(models.StepEvaluate))
	UserActivityModel = Gorm.Model(new(models.UserActivity))
	UserEvaluateModel = Gorm.Model(new(models.UserEvaluate))
	UserPassModel = Gorm.Model(new(models.UserPass))
}
