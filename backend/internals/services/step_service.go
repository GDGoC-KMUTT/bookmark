package services

import (
	"backend/internals/db/models"
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"backend/internals/utils"
	"fmt"
	"sort"
	"strconv"
	"time"
)

type stepService struct {
	stepEvalRepo          repositories.StepEvaluateRepository
	userEvalRepo          repositories.UserEvaluateRepository
	userRepo              repositories.UserRepository
	stepCommentRepo       repositories.StepCommentRepository
	stepCommentUpVoteRepo repositories.StepCommentUpVoteRepository
	stepRepo              repositories.StepRepository
	userPassedRepo        repositories.UserPassedRepository
	stepAuthorRepo        repositories.StepAuthorRepository
	courseContentRepo     repositories.CourseContentRepository
}

func NewStepService(stepEvalRepo repositories.StepEvaluateRepository, userEvalRepo repositories.UserEvaluateRepository, userRepo repositories.UserRepository, stepCommentRepo repositories.StepCommentRepository, stepCommentUpVoteRepo repositories.StepCommentUpVoteRepository, stepRepo repositories.StepRepository, userPassedRepo repositories.UserPassedRepository, stepAuthorRepo repositories.StepAuthorRepository, courseContentRepo repositories.CourseContentRepository) StepService {
	return &stepService{
		stepEvalRepo:          stepEvalRepo,
		userEvalRepo:          userEvalRepo,
		userRepo:              userRepo,
		stepCommentRepo:       stepCommentRepo,
		stepCommentUpVoteRepo: stepCommentUpVoteRepo,
		stepRepo:              stepRepo,
		userPassedRepo:        userPassedRepo,
		stepAuthorRepo:        stepAuthorRepo,
		courseContentRepo:     courseContentRepo,
	}
}

func (r *stepService) GetGems(stepId *uint64, userId *float64) (*int, *int, error) {
	stepEvals, err := r.stepEvalRepo.GetStepEvalByStepId(stepId)
	if err != nil {
		return nil, nil, err
	}

	totalGems := 0
	currentGems := 0
	for _, eval := range stepEvals {
		totalGems += *eval.Gem
		userEvals, err2 := r.userEvalRepo.GetUserEvalByStepEvalIdUserId(eval.Id, userId)
		if err2 != nil {
			return nil, nil, err2
		}

		if userEvals == nil {
			continue
		}

		if *userEvals.Pass == true {
			currentGems += *eval.Gem
		}

	}

	return &totalGems, &currentGems, nil
}

func (r *stepService) GetStepComment(stepId *uint64) ([]payload.StepCommentInfo, error) {
	stepComments, err := r.stepCommentRepo.GetStepCommentByStepId(stepId)
	if err != nil {
		return nil, err
	}

	stepCommentInfo := make([]payload.StepCommentInfo, 0)
	for _, comment := range stepComments {
		user, err := r.userRepo.FindUserByID(utils.Ptr(strconv.FormatUint(*comment.UserId, 10)))
		if err != nil {
			return nil, err
		}

		stepCommentUpVote, err := r.stepCommentUpVoteRepo.GetStepCommentUpVoteByStepCommentId(comment.Id)
		if err != nil {
			return nil, err
		}

		stepCommentInfo = append(stepCommentInfo, payload.StepCommentInfo{
			StepCommentId: comment.Id,
			UserInfo: &payload.CommentedBy{
				UserId:    user.Id,
				FirstName: user.Firstname,
				Lastname:  user.Lastname,
				Email:     user.Email,
				PhotoUrl:  user.PhotoUrl,
			},
			Comment: comment.Content,
			UpVote:  utils.Ptr(len(stepCommentUpVote)),
		})
	}

	return stepCommentInfo, nil
}

func (r *stepService) CreteStpComment(stepId *uint64, userId *float64, content *string) error {
	stepComment := &models.StepComment{
		Content: content,
		StepId:  stepId,
		UserId:  utils.Ptr(uint64(*userId)),
	}

	if err := r.stepCommentRepo.CreateStepComment(stepComment); err != nil {
		return err
	}

	return nil
}

func (r *stepService) CreateStepCommentUpVote(userId *float64, stepCommentId *uint64) error {
	stepCommentUpVote := &models.StepCommentUpvote{
		StepCommentId: stepCommentId,
		UserId:        utils.Ptr(uint64(*userId)),
	}

	if err := r.stepCommentUpVoteRepo.CreateStepCommentUpVote(stepCommentUpVote); err != nil {
		return err
	}

	return nil
}

func (r *stepService) GetStepInfo(stepId *uint64) (*payload.StepInfo, error) {
	step, err := r.stepRepo.GetStepById(stepId)
	if err != nil {
		return nil, err
	}

	stepAuthor, err := r.stepAuthorRepo.GetStepAuthorByStepId(stepId)
	if err != nil {
		return nil, err
	}

	stepEvals, err := r.stepEvalRepo.GetStepEvalByStepId(stepId)
	if err != nil {
		return nil, err
	}

	stepDetail := &payload.StepDetail{
		StepId:      step.Id,
		ModuleId:    step.ModuleId,
		Title:       step.Title,
		Description: step.Description,
		Content:     step.Content,
		Outcome:     step.Outcome,
		Check:       step.Check,
		Error:       step.Error,
	}

	stepInfo := &payload.StepInfo{
		Step: stepDetail,
	}

	authors := make([]*models.User, 0)
	for _, author := range stepAuthor {
		user, err := r.userRepo.FindUserByID(utils.Ptr(strconv.FormatUint(*author.UserId, 10)))
		if err != nil {
			return nil, err
		}
		authors = append(authors, user)
	}
	stepInfo.Authors = authors

	passedUsersMap := make(map[uint64]int)
	for _, stepEval := range stepEvals {
		userEvals, err := r.userEvalRepo.GetPassAllUserEvalByStepEvalId(stepEval.Id)
		if err != nil {
			return nil, err
		}

		// Increment pass count for each user
		for _, userEval := range userEvals {
			userId := *userEval.UserId
			passedUsersMap[userId]++
		}
	}
	fmt.Println(passedUsersMap)

	requiredPassCount := len(stepEvals)
	passedUsers := make([]uint64, 0)
	for userId, passCount := range passedUsersMap {
		if passCount == requiredPassCount {
			passedUsers = append(passedUsers, userId)
		}
	}

	// Fetch user details for the passed users
	users := make([]*models.User, 0, len(passedUsers))
	for _, userId := range passedUsers {
		user, err := r.userRepo.FindUserByID(utils.Ptr(strconv.FormatUint(userId, 10)))
		if err != nil {
			return nil, fmt.Errorf("failed to fetch user by ID: %w", err)
		}
		users = append(users, user)
	}
	stepInfo.UserPassed = users

	return stepInfo, nil
}

func (r *stepService) GetStepEvalInfo(stepId *uint64, userId *float64) ([]*payload.StepEvalInfo, error) {
	stepEvals, err := r.stepEvalRepo.GetStepEvalByStepId(stepId)
	if err != nil {
		return nil, err
	}

	stepEvalInfoList := make([]*payload.StepEvalInfo, 0)

	for _, eval := range stepEvals {
		result := &payload.StepEvalInfo{
			StepId:      eval.StepId,
			StepEvalId:  eval.Id,
			Order:       eval.Order,
			Instruction: eval.Instruction,
			Type:        eval.Type,
			Question:    eval.Question,
		}

		userEval, err := r.userEvalRepo.GetUserEvalByStepEvalIdUserId(eval.Id, userId)
		if err != nil {
			return nil, err
		}

		evalResult := &payload.UserEvalResult{
			Content: userEval.Content,
			Pass:    userEval.Pass,
			Comment: userEval.Comment,
		}
		result.UserEval = evalResult

		stepEvalInfoList = append(stepEvalInfoList, result)
	}

	// order result by Order field
	sort.SliceStable(stepEvalInfoList, func(i, j int) bool {
		return *stepEvalInfoList[i].Order < *stepEvalInfoList[j].Order
	})

	return stepEvalInfoList, nil
}

func (r *stepService) CreateFileFormat(stepId *uint64, stepEvalId *uint64, userId *float64) (*string, error) {
	moduleId, err := r.stepRepo.GetModuleIdByStepId(stepId)
	if err != nil {
		return nil, err
	}

	courseId, err := r.courseContentRepo.GetCourseIdByModuleId(moduleId)
	if err != nil {
		return nil, err
	}

	filename := fmt.Sprintf("course%d_module%d_step%d_userId%d_eval%d_%s.png", *courseId, *moduleId, *stepId, uint64(*userId), *stepEvalId, time.Now().UTC().Format(time.RFC3339))

	return &filename, nil
}

func (r *stepService) CreateUserEval(payload *payload.CreateUserEvalReq) (*uint64, error) {
	userEval := &models.UserEvaluate{
		UserId:         utils.Ptr(uint64(*payload.UserId)),
		Content:        payload.Content,
		StepEvaluateId: payload.StepEvalId,
	}

	result, err := r.userEvalRepo.CreateUserEval(userEval)
	if err != nil {
		return nil, err
	}

	return result.Id, nil
}

func (r *stepService) CheckStepEvalStatus(userEvalId *uint64) (*models.UserEvaluate, error) {
	userEval, err := r.userEvalRepo.GetUserEvalById(userEvalId)
	if err != nil {
		return nil, err
	}

	if userEval.Pass == nil && userEval.Comment == nil {
		return nil, nil
	}

	// TODO
	//stepEval, err := r.stepEvalRepo.GetStepEvalById(userEval.StepEvaluateId)
	//if err != nil {
	//	return nil, err
	//}

	return userEval, nil
}