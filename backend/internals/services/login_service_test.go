package services_test

import (
	"backend/internals/config"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LoginServiceTestSuite struct {
	suite.Suite
}

func TestLoginService(t *testing.T) {
	config.BootConfiguration()
	suite.Run(t, new(LoginServiceTestSuite))
}
