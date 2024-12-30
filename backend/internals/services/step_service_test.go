package services

import (
	"backend/internals/db/models"
	"backend/internals/entities/payload"
	"backend/internals/utils"
	mockRepositories "backend/mocks/repositories"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
)

type StepServiceTestSuite struct {
	suite.Suite
}

func (suite *StepServiceTestSuite) TestGetGemsWhenSuccess() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(2))
	mockUserId := utils.Ptr(float64(1))
	mockStepEval := []*models.StepEvaluate{
		{
			StepId: mockStepId,
			Id:     utils.Ptr(uint64(1)),
			Gem:    utils.Ptr(2),
		},
	}
	mockUserEval := &models.UserEvaluate{
		Pass: utils.Ptr(true),
	}

	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(mockStepEval, nil)
	mockUserEvalRepo.EXPECT().GetUserEvalByStepEvalIdUserId(mock.Anything, mock.Anything).Return(mockUserEval, nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	totalGem, currentGem, err := underTest.GetGems(mockStepId, mockUserId)

	is.Nil(err)
	is.Equal(2, *totalGem)
	is.Equal(2, *currentGem)
}

func (suite *StepServiceTestSuite) TestGetGemsWhenPassNil() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(2))
	mockUserId := utils.Ptr(float64(1))
	mockStepEval := []*models.StepEvaluate{
		{
			StepId: mockStepId,
			Id:     utils.Ptr(uint64(1)),
			Gem:    utils.Ptr(2),
		},
	}
	mockUserEval := &models.UserEvaluate{}

	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(mockStepEval, nil)
	mockUserEvalRepo.EXPECT().GetUserEvalByStepEvalIdUserId(mock.Anything, mock.Anything).Return(mockUserEval, nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	totalGem, currentGem, err := underTest.GetGems(mockStepId, mockUserId)

	is.Nil(err)
	is.Equal(2, *totalGem)
	is.Equal(0, *currentGem)
}

func (suite *StepServiceTestSuite) TestGetGemsWhenUserEvalNil() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(2))
	mockUserId := utils.Ptr(float64(1))
	mockStepEval := []*models.StepEvaluate{
		{
			StepId: mockStepId,
			Id:     utils.Ptr(uint64(1)),
			Gem:    utils.Ptr(2),
		},
	}

	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(mockStepEval, nil)
	mockUserEvalRepo.EXPECT().GetUserEvalByStepEvalIdUserId(mock.Anything, mock.Anything).Return(nil, nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	totalGem, currentGem, err := underTest.GetGems(mockStepId, mockUserId)

	is.Nil(err)
	is.Equal(2, *totalGem)
	is.Equal(0, *currentGem)
}
func (suite *StepServiceTestSuite) TestGetGemsWhenFailedToGetStepEvalByStepId() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(2))
	mockUserId := utils.Ptr(float64(1))

	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(nil, fmt.Errorf("failed to getStepEval"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	totalGem, currentGem, err := underTest.GetGems(mockStepId, mockUserId)

	is.NotNil(err)
	is.Nil(totalGem)
	is.Nil(currentGem)
	is.Equal("failed to getStepEval", err.Error())
}

func (suite *StepServiceTestSuite) TestGetGemsWhenFailedToGetUserEval() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(2))
	mockUserId := utils.Ptr(float64(1))
	mockStepEval := []*models.StepEvaluate{
		{
			StepId: mockStepId,
			Id:     utils.Ptr(uint64(1)),
			Gem:    utils.Ptr(2),
		},
	}
	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(mockStepEval, nil)
	mockUserEvalRepo.EXPECT().GetUserEvalByStepEvalIdUserId(mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to getUserEvalByStepEvalIdUserId"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	totalGem, currentGem, err := underTest.GetGems(mockStepId, mockUserId)

	is.NotNil(err)
	is.Nil(totalGem)
	is.Nil(currentGem)
	is.Equal("failed to getUserEvalByStepEvalIdUserId", err.Error())
}

func (suite *StepServiceTestSuite) TestGetStepCommentWhenSuccess() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(2))
	mockUserId := utils.Ptr(float64(1))
	mockStepComments := []*models.StepComment{
		{
			Id:     utils.Ptr(uint64(1)),
			StepId: mockStepId,
			UserId: utils.Ptr(uint64(*mockUserId)),
		},
	}
	mockUser := &models.User{
		Id: utils.Ptr(uint64(*mockUserId)),
	}

	mockStepCommentUpVote := []*models.StepCommentUpvote{
		{
			UserId: utils.Ptr(uint64(*mockUserId)),
		},
	}

	mockStepCommentRepo.EXPECT().GetStepCommentByStepId(mock.Anything).Return(mockStepComments, nil)
	mockUserRepo.EXPECT().FindUserByID(mock.Anything).Return(mockUser, nil)
	mockStepCommentUpVoteRepo.EXPECT().GetStepCommentUpVoteByStepCommentId(mock.Anything).Return(mockStepCommentUpVote, nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepCommentInfo, err := underTest.GetStepComment(mockStepId, utils.Ptr(uint64(*mockUserId)))

	is.Nil(err)
	is.Equal(true, *stepCommentInfo[0].HasUpVoted)
}

func (suite *StepServiceTestSuite) TestGetStepCommentWhenFailedToGetStepCommentByStepId() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(2))
	mockUserId := utils.Ptr(float64(1))

	mockStepCommentRepo.EXPECT().GetStepCommentByStepId(mock.Anything).Return(nil, fmt.Errorf("failed to get stepComment by stepId"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepCommentInfo, err := underTest.GetStepComment(mockStepId, utils.Ptr(uint64(*mockUserId)))

	is.NotNil(err)
	is.Nil(stepCommentInfo)
	is.Equal("failed to get stepComment by stepId", err.Error())
}

func (suite *StepServiceTestSuite) TestGetStepCommentWhenFailedToFindUserById() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(2))
	mockUserId := utils.Ptr(float64(1))
	mockStepComments := []*models.StepComment{
		{
			Id:     utils.Ptr(uint64(1)),
			StepId: mockStepId,
			UserId: utils.Ptr(uint64(*mockUserId)),
		},
	}

	mockStepCommentRepo.EXPECT().GetStepCommentByStepId(mock.Anything).Return(mockStepComments, nil)
	mockUserRepo.EXPECT().FindUserByID(mock.Anything).Return(nil, fmt.Errorf("failed to find user by id"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepCommentInfo, err := underTest.GetStepComment(mockStepId, utils.Ptr(uint64(*mockUserId)))

	is.NotNil(err)
	is.Nil(stepCommentInfo)
	is.Equal("failed to find user by id", err.Error())
}

func (suite *StepServiceTestSuite) TestGetStepCommentWhenFailedToGetStepCommentUpVote() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(2))
	mockUserId := utils.Ptr(float64(1))
	mockStepComments := []*models.StepComment{
		{
			Id:     utils.Ptr(uint64(1)),
			StepId: mockStepId,
			UserId: utils.Ptr(uint64(*mockUserId)),
		},
	}
	mockUser := &models.User{
		Id: utils.Ptr(uint64(*mockUserId)),
	}

	mockStepCommentRepo.EXPECT().GetStepCommentByStepId(mock.Anything).Return(mockStepComments, nil)
	mockUserRepo.EXPECT().FindUserByID(mock.Anything).Return(mockUser, nil)
	mockStepCommentUpVoteRepo.EXPECT().GetStepCommentUpVoteByStepCommentId(mock.Anything).Return(nil, fmt.Errorf("failed to get stepCommentUpvote"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepCommentInfo, err := underTest.GetStepComment(mockStepId, utils.Ptr(uint64(*mockUserId)))

	is.NotNil(err)
	is.Nil(stepCommentInfo)
	is.Equal("failed to get stepCommentUpvote", err.Error())
}
func (suite *StepServiceTestSuite) TestCreateStepCommentWhenSuccess() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(2))
	mockUserId := utils.Ptr(float64(1))
	mockContent := utils.Ptr("comment")

	mockStepCommentRepo.EXPECT().CreateStepComment(mock.Anything).Return(nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	err := underTest.CreateStpComment(mockStepId, mockUserId, mockContent)

	is.Nil(err)
}

func (suite *StepServiceTestSuite) TestCreateStepCommentWhenFailedToCreateComment() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(2))
	mockUserId := utils.Ptr(float64(1))
	mockContent := utils.Ptr("comment")

	mockStepCommentRepo.EXPECT().CreateStepComment(mock.Anything).Return(fmt.Errorf("failed to create comment"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	err := underTest.CreateStpComment(mockStepId, mockUserId, mockContent)

	is.NotNil(err)
	is.Equal("failed to create comment", err.Error())
}

func (suite *StepServiceTestSuite) TestCreateOrDeleteStepCommentUpVoteWhenNoExitUpVote() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserId := utils.Ptr(float64(1))
	mockStepCommentId := utils.Ptr(uint64(1))

	mockStepCommentUpVoteRepo.EXPECT().GetStepCommentUpVoteByStepCommentIdAndUserId(mock.Anything, mock.Anything).Return(nil, nil)
	mockStepCommentUpVoteRepo.EXPECT().CreateStepCommentUpVote(mock.Anything).Return(nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	err := underTest.CreateOrDeleteStepCommentUpVote(mockUserId, mockStepCommentId)

	is.Nil(err)
}

func (suite *StepServiceTestSuite) TestCreateOrDeleteStepCommentUpVoteWhenNoExitUpVoteAndFailedToGetExistStepComment() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserId := utils.Ptr(float64(1))
	mockStepCommentId := utils.Ptr(uint64(1))

	mockStepCommentUpVoteRepo.EXPECT().GetStepCommentUpVoteByStepCommentIdAndUserId(mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to get stepCommentUpVote"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	err := underTest.CreateOrDeleteStepCommentUpVote(mockUserId, mockStepCommentId)

	is.NotNil(err)
	is.Equal("failed to get stepCommentUpVote", err.Error())
}

func (suite *StepServiceTestSuite) TestCreateOrDeleteStepCommentUpVoteWhenNoExitUpVoteAndFailedToCreateComment() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserId := utils.Ptr(float64(1))
	mockStepCommentId := utils.Ptr(uint64(1))

	mockStepCommentUpVoteRepo.EXPECT().GetStepCommentUpVoteByStepCommentIdAndUserId(mock.Anything, mock.Anything).Return(nil, nil)
	mockStepCommentUpVoteRepo.EXPECT().CreateStepCommentUpVote(mock.Anything).Return(fmt.Errorf("failed to create comment"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	err := underTest.CreateOrDeleteStepCommentUpVote(mockUserId, mockStepCommentId)

	is.NotNil(err)
	is.Equal("failed to create comment", err.Error())
}

func (suite *StepServiceTestSuite) TestCreateOrDeleteStepCommentUpVoteWhenHaveExitUpVote() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserId := utils.Ptr(float64(1))
	mockStepCommentId := utils.Ptr(uint64(1))
	mockStepCommentUpVote := &models.StepCommentUpvote{
		StepCommentId: mockStepCommentId,
	}

	mockStepCommentUpVoteRepo.EXPECT().GetStepCommentUpVoteByStepCommentIdAndUserId(mock.Anything, mock.Anything).Return(mockStepCommentUpVote, nil)
	mockStepCommentUpVoteRepo.EXPECT().DeleteStepCommentUpVote(mock.Anything, mock.Anything).Return(nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	err := underTest.CreateOrDeleteStepCommentUpVote(mockUserId, mockStepCommentId)

	is.Nil(err)
}

func (suite *StepServiceTestSuite) TestCreateOrDeleteStepCommentUpVoteWhenHaveExitUpVoteAndFailedToDeleteComment() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserId := utils.Ptr(float64(1))
	mockStepCommentId := utils.Ptr(uint64(1))
	mockStepCommentUpVote := &models.StepCommentUpvote{
		StepCommentId: mockStepCommentId,
	}

	mockStepCommentUpVoteRepo.EXPECT().GetStepCommentUpVoteByStepCommentIdAndUserId(mock.Anything, mock.Anything).Return(mockStepCommentUpVote, nil)
	mockStepCommentUpVoteRepo.EXPECT().DeleteStepCommentUpVote(mock.Anything, mock.Anything).Return(fmt.Errorf("failed to delete comment"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	err := underTest.CreateOrDeleteStepCommentUpVote(mockUserId, mockStepCommentId)

	is.NotNil(err)
	is.Equal("failed to delete comment", err.Error())
}

func (suite *StepServiceTestSuite) TestGetStepEvalInfoTypeTextWhenSuccess() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserId := utils.Ptr(float64(1))
	mockStepId := utils.Ptr(uint64(1))

	mockStepEvalId := utils.Ptr(uint64(1))
	mockOrder := utils.Ptr(1)
	mockInstruction := utils.Ptr("intrs")
	mockType := utils.Ptr("text")
	mockQuestion := utils.Ptr("question")

	mockStepEvals := []*models.StepEvaluate{
		{
			StepId:      mockStepId,
			Id:          mockStepEvalId,
			Order:       mockOrder,
			Instruction: mockInstruction,
			Type:        mockType,
			Question:    mockQuestion,
		},
	}

	mockUserEval := &models.UserEvaluate{
		Id:      utils.Ptr(uint64(1)),
		Content: utils.Ptr("content"),
		Pass:    utils.Ptr(true),
		Comment: utils.Ptr("comment"),
	}

	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(mockStepEvals, nil)
	mockUserEvalRepo.EXPECT().GetUserEvalByStepEvalIdUserId(mock.Anything, mock.Anything).Return(mockUserEval, nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepEvals, err := underTest.GetStepEvalInfo(mockStepId, mockUserId)

	is.Nil(err)
	is.NotNil(stepEvals)
}

func (suite *StepServiceTestSuite) TestGetStepEvalInfoTypeImageWhenSuccess() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserId := utils.Ptr(float64(1))
	mockStepId := utils.Ptr(uint64(1))

	mockStepEvalId := utils.Ptr(uint64(1))
	mockOrder := utils.Ptr(1)
	mockInstruction := utils.Ptr("intrs")
	mockType := utils.Ptr("image")
	mockQuestion := utils.Ptr("question")

	mockStepEvals := []*models.StepEvaluate{
		{
			StepId:      mockStepId,
			Id:          mockStepEvalId,
			Order:       mockOrder,
			Instruction: mockInstruction,
			Type:        mockType,
			Question:    mockQuestion,
		},
	}

	mockUserEval := &models.UserEvaluate{
		Id:      utils.Ptr(uint64(1)),
		Content: utils.Ptr("content"),
		Pass:    utils.Ptr(true),
		Comment: utils.Ptr("comment"),
	}

	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(mockStepEvals, nil)
	mockUserEvalRepo.EXPECT().GetUserEvalByStepEvalIdUserId(mock.Anything, mock.Anything).Return(mockUserEval, nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepEvals, err := underTest.GetStepEvalInfo(mockStepId, mockUserId)

	is.Nil(err)
	is.NotNil(stepEvals)
}

func (suite *StepServiceTestSuite) TestGetStepEvalInfoWhenFailedToGetStepEval() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserId := utils.Ptr(float64(1))
	mockStepId := utils.Ptr(uint64(1))

	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(nil, fmt.Errorf("failed to get step eval"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepEvals, err := underTest.GetStepEvalInfo(mockStepId, mockUserId)

	is.NotNil(err)
	is.Nil(stepEvals)
	is.Equal("failed to get step eval", err.Error())
}

func (suite *StepServiceTestSuite) TestGetStepEvalInfoTypeImageWhenFailedToGetUserEval() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserId := utils.Ptr(float64(1))
	mockStepId := utils.Ptr(uint64(1))

	mockStepEvalId := utils.Ptr(uint64(1))
	mockOrder := utils.Ptr(1)
	mockInstruction := utils.Ptr("intrs")
	mockType := utils.Ptr("image")
	mockQuestion := utils.Ptr("question")

	mockStepEvals := []*models.StepEvaluate{
		{
			StepId:      mockStepId,
			Id:          mockStepEvalId,
			Order:       mockOrder,
			Instruction: mockInstruction,
			Type:        mockType,
			Question:    mockQuestion,
		},
	}

	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(mockStepEvals, nil)
	mockUserEvalRepo.EXPECT().GetUserEvalByStepEvalIdUserId(mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to get user eval"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepEvals, err := underTest.GetStepEvalInfo(mockStepId, mockUserId)

	is.NotNil(err)
	is.Nil(stepEvals)
	is.Equal("failed to get user eval", err.Error())
}

func (suite *StepServiceTestSuite) TestCreateFileFormatWhenSuccess() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserId := utils.Ptr(float64(1))
	mockStepId := utils.Ptr(uint64(1))
	mockStepEvalId := utils.Ptr(uint64(1))
	mockModuleId := utils.Ptr(uint64(2))
	mockCourseId := utils.Ptr(uint64(1))

	mockStepRepo.EXPECT().GetModuleIdByStepId(mock.Anything).Return(mockModuleId, nil)
	mockCourseContentRepo.EXPECT().GetCourseIdByModuleId(mock.Anything).Return(mockCourseId, nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	filename, err := underTest.CreateFileFormat(mockStepId, mockStepEvalId, mockUserId)

	is.Nil(err)
	is.NotNil(filename)
}

func (suite *StepServiceTestSuite) TestCreateFileFormatWhenFailedToGetModule() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserId := utils.Ptr(float64(1))
	mockStepId := utils.Ptr(uint64(1))
	mockStepEvalId := utils.Ptr(uint64(1))

	mockStepRepo.EXPECT().GetModuleIdByStepId(mock.Anything).Return(nil, fmt.Errorf("failed to get moduleId"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	filename, err := underTest.CreateFileFormat(mockStepId, mockStepEvalId, mockUserId)

	is.NotNil(err)
	is.Nil(filename)
	is.Equal("failed to get moduleId", err.Error())
}

func (suite *StepServiceTestSuite) TestCreateFileFormatWhenFailedToGetCourse() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserId := utils.Ptr(float64(1))
	mockStepId := utils.Ptr(uint64(1))
	mockStepEvalId := utils.Ptr(uint64(1))
	mockModuleId := utils.Ptr(uint64(2))

	mockStepRepo.EXPECT().GetModuleIdByStepId(mock.Anything).Return(mockModuleId, nil)
	mockCourseContentRepo.EXPECT().GetCourseIdByModuleId(mock.Anything).Return(nil, fmt.Errorf("failed to get courseId"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	filename, err := underTest.CreateFileFormat(mockStepId, mockStepEvalId, mockUserId)

	is.NotNil(err)
	is.Nil(filename)
	is.Equal("failed to get courseId", err.Error())
}

func (suite *StepServiceTestSuite) TestCreateUserEvalWhenSuccess() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockPayload := &payload.CreateUserEvalReq{
		UserId:     utils.Ptr(float64(1)),
		Content:    utils.Ptr("content"),
		StepEvalId: utils.Ptr(uint64(1)),
	}

	mockCreatedUserEval := &models.UserEvaluate{
		Id: utils.Ptr(uint64(12)),
	}

	mockUserEvalRepo.EXPECT().CreateUserEval(mock.Anything).Return(mockCreatedUserEval, nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	userEvalId, err := underTest.CreateUserEval(mockPayload)

	is.Nil(err)
	is.NotNil(userEvalId)
	is.Equal(uint64(12), *userEvalId)
}

func (suite *StepServiceTestSuite) TestCreateUserEvalWhenFailedToCreateUserEval() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockPayload := &payload.CreateUserEvalReq{
		UserId:     utils.Ptr(float64(1)),
		Content:    utils.Ptr("content"),
		StepEvalId: utils.Ptr(uint64(1)),
	}

	mockUserEvalRepo.EXPECT().CreateUserEval(mock.Anything).Return(nil, fmt.Errorf("failed to create user eval"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	userEvalId, err := underTest.CreateUserEval(mockPayload)

	is.NotNil(err)
	is.Nil(userEvalId)
	is.Equal("failed to create user eval", err.Error())
}

func (suite *StepServiceTestSuite) TestCheckStepEvalStatusWhenSuccess() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserEvalId := utils.Ptr(uint64(12))
	mockUserId := utils.Ptr(uint64(1))
	mockUserEval := &models.UserEvaluate{
		Pass:    utils.Ptr(true),
		Comment: utils.Ptr("comment"),
		Id:      utils.Ptr(uint64(1)),
		Content: utils.Ptr("content"),
	}

	mockUserEvalRepo.EXPECT().GetUserEvalByIdAndUserId(mock.Anything, mock.Anything).Return(mockUserEval, nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	userEvalResult, err := underTest.CheckStepEvalStatus(mockUserEvalId, mockUserId)

	is.Nil(err)
	is.NotNil(userEvalResult)
}

func (suite *StepServiceTestSuite) TestCheckStepEvalStatusWhenFailedToGetUserEval() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserEvalId := utils.Ptr(uint64(12))
	mockUserId := utils.Ptr(uint64(1))

	mockUserEvalRepo.EXPECT().GetUserEvalByIdAndUserId(mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to get user eval"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	userEvalResult, err := underTest.CheckStepEvalStatus(mockUserEvalId, mockUserId)

	is.NotNil(err)
	is.Nil(userEvalResult)
	is.Equal("failed to get user eval", err.Error())
}

func (suite *StepServiceTestSuite) TestCheckStepEvalStatusWhenUserEvalNil() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserEvalId := utils.Ptr(uint64(12))
	mockUserId := utils.Ptr(uint64(1))

	mockUserEvalRepo.EXPECT().GetUserEvalByIdAndUserId(mock.Anything, mock.Anything).Return(nil, nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	userEvalResult, err := underTest.CheckStepEvalStatus(mockUserEvalId, mockUserId)

	is.Nil(err)
	is.Nil(userEvalResult)
}

func (suite *StepServiceTestSuite) TestCheckStepEvalStatusWhenPassNilCommentNil() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockUserEvalId := utils.Ptr(uint64(12))
	mockUserId := utils.Ptr(uint64(1))
	mockUserEval := &models.UserEvaluate{
		Id:      utils.Ptr(uint64(1)),
		Content: utils.Ptr("content"),
	}

	mockUserEvalRepo.EXPECT().GetUserEvalByIdAndUserId(mock.Anything, mock.Anything).Return(mockUserEval, nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	userEvalResult, err := underTest.CheckStepEvalStatus(mockUserEvalId, mockUserId)

	is.Nil(err)
	is.Nil(userEvalResult)
}

func (suite *StepServiceTestSuite) TestSubmitStepEvalTypeCheckWhenSuccess() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepEvalId := utils.Ptr(uint64(12))
	mockUserId := utils.Ptr(uint64(1))
	mockUserEval := &models.UserEvaluate{
		Id: utils.Ptr(uint64(1)),
	}

	mockUserEvalRepo.EXPECT().CreateUserEval(mock.Anything).Return(mockUserEval, nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	userEvalId, err := underTest.SubmitStepEvalTypeCheck(mockStepEvalId, mockUserId)

	is.Nil(err)
	is.NotNil(userEvalId)
	is.Equal(uint64(1), *userEvalId)
}

func (suite *StepServiceTestSuite) TestSubmitStepEvalTypeCheckWhenFailedToCreateUserEval() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepEvalId := utils.Ptr(uint64(12))
	mockUserId := utils.Ptr(uint64(1))

	mockUserEvalRepo.EXPECT().CreateUserEval(mock.Anything).Return(nil, fmt.Errorf("failed to create user eval"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	userEvalId, err := underTest.SubmitStepEvalTypeCheck(mockStepEvalId, mockUserId)

	is.NotNil(err)
	is.Nil(userEvalId)
	is.Equal("failed to create user eval", err.Error())
}

func (suite *StepServiceTestSuite) TestGetStepInfoWhenSuccess() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(1))
	mockUserIdPassed := utils.Ptr(uint64(9))
	mockAuthorId := utils.Ptr(uint64(12))
	mockStep := &models.Step{
		Id:          mockStepId,
		ModuleId:    utils.Ptr(uint64(1)),
		Title:       utils.Ptr("title"),
		Description: utils.Ptr("desc"),
		Content:     utils.Ptr("content"),
		Outcome:     utils.Ptr("outcome"),
		Check:       utils.Ptr("check"),
		Error:       utils.Ptr("error"),
	}
	mockModule := &models.Module{
		ImageUrl: utils.Ptr("img"),
	}
	mockStepAuthors := []*models.StepAuthor{
		{
			UserId: mockAuthorId,
		},
	}
	mockStepEvals := []*models.StepEvaluate{
		{
			Id: utils.Ptr(uint64(12)),
		},
	}
	mockAuthorUser := &models.User{
		Id:        mockAuthorId,
		Firstname: utils.Ptr("fn"),
		Lastname:  utils.Ptr("ln"),
		Email:     utils.Ptr("email"),
		PhotoUrl:  utils.Ptr("photo"),
	}
	mockUserPass := &models.User{
		Id:        mockUserIdPassed,
		Firstname: utils.Ptr("fn"),
		Lastname:  utils.Ptr("ln"),
		Email:     utils.Ptr("email"),
		PhotoUrl:  utils.Ptr("photo"),
	}

	mockUserEval := []*models.UserEvaluate{
		{
			UserId: mockUserIdPassed,
		},
	}

	mockStepRepo.EXPECT().GetStepById(mock.Anything).Return(mockStep, nil)
	mockModuleRepo.EXPECT().GetModuleById(mock.Anything).Return(mockModule, nil)
	mockStepAuthorRepo.EXPECT().GetStepAuthorByStepId(mock.Anything).Return(mockStepAuthors, nil)
	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(mockStepEvals, nil)
	mockUserRepo.EXPECT().FindUserByID(utils.Ptr(strconv.FormatUint(*mockAuthorId, 10))).Return(mockAuthorUser, nil)
	mockUserEvalRepo.EXPECT().GetPassAllUserEvalByStepEvalId(mock.Anything).Return(mockUserEval, nil)
	mockUserRepo.EXPECT().FindUserByID(utils.Ptr(strconv.FormatUint(*mockUserIdPassed, 10))).Return(mockUserPass, nil)

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepInfo, err := underTest.GetStepInfo(mockStepId)

	is.Nil(err)
	is.NotNil(stepInfo)
}

func (suite *StepServiceTestSuite) TestGetStepInfoWhenFailedToGetStep() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(1))

	mockStepRepo.EXPECT().GetStepById(mock.Anything).Return(nil, fmt.Errorf("failed to get step"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepInfo, err := underTest.GetStepInfo(mockStepId)

	is.NotNil(err)
	is.Nil(stepInfo)
	is.Equal("failed to get step", err.Error())
}

func (suite *StepServiceTestSuite) TestGetStepInfoWhenFailedToGetModule() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(1))
	mockStep := &models.Step{
		Id:          mockStepId,
		ModuleId:    utils.Ptr(uint64(1)),
		Title:       utils.Ptr("title"),
		Description: utils.Ptr("desc"),
		Content:     utils.Ptr("content"),
		Outcome:     utils.Ptr("outcome"),
		Check:       utils.Ptr("check"),
		Error:       utils.Ptr("error"),
	}

	mockStepRepo.EXPECT().GetStepById(mock.Anything).Return(mockStep, nil)
	mockModuleRepo.EXPECT().GetModuleById(mock.Anything).Return(nil, fmt.Errorf("failed to get module"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepInfo, err := underTest.GetStepInfo(mockStepId)

	is.NotNil(err)
	is.Nil(stepInfo)
	is.Equal("failed to get module", err.Error())
}

func (suite *StepServiceTestSuite) TestGetStepInfoWhenFailedToGetStepAuthors() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(1))
	mockStep := &models.Step{
		Id:          mockStepId,
		ModuleId:    utils.Ptr(uint64(1)),
		Title:       utils.Ptr("title"),
		Description: utils.Ptr("desc"),
		Content:     utils.Ptr("content"),
		Outcome:     utils.Ptr("outcome"),
		Check:       utils.Ptr("check"),
		Error:       utils.Ptr("error"),
	}
	mockModule := &models.Module{
		ImageUrl: utils.Ptr("img"),
	}

	mockStepRepo.EXPECT().GetStepById(mock.Anything).Return(mockStep, nil)
	mockModuleRepo.EXPECT().GetModuleById(mock.Anything).Return(mockModule, nil)
	mockStepAuthorRepo.EXPECT().GetStepAuthorByStepId(mock.Anything).Return(nil, fmt.Errorf("failed to get step authors"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepInfo, err := underTest.GetStepInfo(mockStepId)

	is.NotNil(err)
	is.Nil(stepInfo)
	is.Equal("failed to get step authors", err.Error())
}
func (suite *StepServiceTestSuite) TestGetStepInfoWhenFailedToGetStepEval() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(1))
	mockAuthorId := utils.Ptr(uint64(12))
	mockStep := &models.Step{
		Id:          mockStepId,
		ModuleId:    utils.Ptr(uint64(1)),
		Title:       utils.Ptr("title"),
		Description: utils.Ptr("desc"),
		Content:     utils.Ptr("content"),
		Outcome:     utils.Ptr("outcome"),
		Check:       utils.Ptr("check"),
		Error:       utils.Ptr("error"),
	}
	mockModule := &models.Module{
		ImageUrl: utils.Ptr("img"),
	}
	mockStepAuthors := []*models.StepAuthor{
		{
			UserId: mockAuthorId,
		},
	}

	mockStepRepo.EXPECT().GetStepById(mock.Anything).Return(mockStep, nil)
	mockModuleRepo.EXPECT().GetModuleById(mock.Anything).Return(mockModule, nil)
	mockStepAuthorRepo.EXPECT().GetStepAuthorByStepId(mock.Anything).Return(mockStepAuthors, nil)
	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(nil, fmt.Errorf("failed to get step eval"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepInfo, err := underTest.GetStepInfo(mockStepId)

	is.NotNil(err)
	is.Nil(stepInfo)
	is.Equal("failed to get step eval", err.Error())
}
func (suite *StepServiceTestSuite) TestGetStepInfoWhenFailedToFindAuthorInfo() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(1))
	mockAuthorId := utils.Ptr(uint64(12))
	mockStep := &models.Step{
		Id:          mockStepId,
		ModuleId:    utils.Ptr(uint64(1)),
		Title:       utils.Ptr("title"),
		Description: utils.Ptr("desc"),
		Content:     utils.Ptr("content"),
		Outcome:     utils.Ptr("outcome"),
		Check:       utils.Ptr("check"),
		Error:       utils.Ptr("error"),
	}
	mockModule := &models.Module{
		ImageUrl: utils.Ptr("img"),
	}
	mockStepAuthors := []*models.StepAuthor{
		{
			UserId: mockAuthorId,
		},
	}
	mockStepEvals := []*models.StepEvaluate{
		{
			Id: utils.Ptr(uint64(12)),
		},
	}

	mockStepRepo.EXPECT().GetStepById(mock.Anything).Return(mockStep, nil)
	mockModuleRepo.EXPECT().GetModuleById(mock.Anything).Return(mockModule, nil)
	mockStepAuthorRepo.EXPECT().GetStepAuthorByStepId(mock.Anything).Return(mockStepAuthors, nil)
	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(mockStepEvals, nil)
	mockUserRepo.EXPECT().FindUserByID(utils.Ptr(strconv.FormatUint(*mockAuthorId, 10))).Return(nil, fmt.Errorf("failed to find author info"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepInfo, err := underTest.GetStepInfo(mockStepId)

	is.NotNil(err)
	is.Nil(stepInfo)
	is.Equal("failed to find author info", err.Error())
}
func (suite *StepServiceTestSuite) TestGetStepInfoWhenFailedToGetPassAllUserEval() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(1))
	mockAuthorId := utils.Ptr(uint64(12))
	mockStep := &models.Step{
		Id:          mockStepId,
		ModuleId:    utils.Ptr(uint64(1)),
		Title:       utils.Ptr("title"),
		Description: utils.Ptr("desc"),
		Content:     utils.Ptr("content"),
		Outcome:     utils.Ptr("outcome"),
		Check:       utils.Ptr("check"),
		Error:       utils.Ptr("error"),
	}
	mockModule := &models.Module{
		ImageUrl: utils.Ptr("img"),
	}
	mockStepAuthors := []*models.StepAuthor{
		{
			UserId: mockAuthorId,
		},
	}
	mockStepEvals := []*models.StepEvaluate{
		{
			Id: utils.Ptr(uint64(12)),
		},
	}
	mockAuthorUser := &models.User{
		Id:        mockAuthorId,
		Firstname: utils.Ptr("fn"),
		Lastname:  utils.Ptr("ln"),
		Email:     utils.Ptr("email"),
		PhotoUrl:  utils.Ptr("photo"),
	}

	mockStepRepo.EXPECT().GetStepById(mock.Anything).Return(mockStep, nil)
	mockModuleRepo.EXPECT().GetModuleById(mock.Anything).Return(mockModule, nil)
	mockStepAuthorRepo.EXPECT().GetStepAuthorByStepId(mock.Anything).Return(mockStepAuthors, nil)
	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(mockStepEvals, nil)
	mockUserRepo.EXPECT().FindUserByID(utils.Ptr(strconv.FormatUint(*mockAuthorId, 10))).Return(mockAuthorUser, nil)
	mockUserEvalRepo.EXPECT().GetPassAllUserEvalByStepEvalId(mock.Anything).Return(nil, fmt.Errorf("failed to get user eval that pass all step eval"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepInfo, err := underTest.GetStepInfo(mockStepId)

	is.NotNil(err)
	is.Nil(stepInfo)
	is.Equal("failed to get user eval that pass all step eval", err.Error())
}
func (suite *StepServiceTestSuite) TestGetStepInfoWhenFailedToFindUserPassedInfo() {
	is := assert.New(suite.T())

	mockStepRepo := new(mockRepositories.StepRepository)
	mockStepEvalRepo := new(mockRepositories.StepEvaluateRepository)
	mockStepCommentRepo := new(mockRepositories.StepCommentRepository)
	mockStepCommentUpVoteRepo := new(mockRepositories.StepCommentUpVoteRepository)
	mockStepAuthorRepo := new(mockRepositories.StepAuthorRepository)

	mockUserRepo := new(mockRepositories.UserRepository)
	mockUserEvalRepo := new(mockRepositories.UserEvaluateRepository)

	mockCourseContentRepo := new(mockRepositories.CourseContentRepository)

	mockModuleRepo := new(mockRepositories.ModulesRepository)

	mockStepId := utils.Ptr(uint64(1))
	mockUserIdPassed := utils.Ptr(uint64(9))
	mockAuthorId := utils.Ptr(uint64(12))
	mockStep := &models.Step{
		Id:          mockStepId,
		ModuleId:    utils.Ptr(uint64(1)),
		Title:       utils.Ptr("title"),
		Description: utils.Ptr("desc"),
		Content:     utils.Ptr("content"),
		Outcome:     utils.Ptr("outcome"),
		Check:       utils.Ptr("check"),
		Error:       utils.Ptr("error"),
	}
	mockModule := &models.Module{
		ImageUrl: utils.Ptr("img"),
	}
	mockStepAuthors := []*models.StepAuthor{
		{
			UserId: mockAuthorId,
		},
	}
	mockStepEvals := []*models.StepEvaluate{
		{
			Id: utils.Ptr(uint64(12)),
		},
	}
	mockAuthorUser := &models.User{
		Id:        mockAuthorId,
		Firstname: utils.Ptr("fn"),
		Lastname:  utils.Ptr("ln"),
		Email:     utils.Ptr("email"),
		PhotoUrl:  utils.Ptr("photo"),
	}

	mockUserEval := []*models.UserEvaluate{
		{
			UserId: mockUserIdPassed,
		},
	}

	mockStepRepo.EXPECT().GetStepById(mock.Anything).Return(mockStep, nil)
	mockModuleRepo.EXPECT().GetModuleById(mock.Anything).Return(mockModule, nil)
	mockStepAuthorRepo.EXPECT().GetStepAuthorByStepId(mock.Anything).Return(mockStepAuthors, nil)
	mockStepEvalRepo.EXPECT().GetStepEvalByStepId(mock.Anything).Return(mockStepEvals, nil)
	mockUserRepo.EXPECT().FindUserByID(utils.Ptr(strconv.FormatUint(*mockAuthorId, 10))).Return(mockAuthorUser, nil)
	mockUserEvalRepo.EXPECT().GetPassAllUserEvalByStepEvalId(mock.Anything).Return(mockUserEval, nil)
	mockUserRepo.EXPECT().FindUserByID(utils.Ptr(strconv.FormatUint(*mockUserIdPassed, 10))).Return(nil, fmt.Errorf("failed to find user passed info"))

	underTest := NewStepService(mockStepRepo, mockStepEvalRepo, mockStepCommentRepo, mockStepCommentUpVoteRepo, mockStepAuthorRepo, mockUserRepo, mockUserEvalRepo, mockCourseContentRepo, mockModuleRepo)

	stepInfo, err := underTest.GetStepInfo(mockStepId)

	is.NotNil(err)
	is.Nil(stepInfo)
	is.Equal("failed to find user passed info", err.Error())
}

func TestStepService(t *testing.T) {
	suite.Run(t, new(StepServiceTestSuite))
}
