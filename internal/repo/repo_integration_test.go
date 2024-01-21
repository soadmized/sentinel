package repo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	repo Repo
}

func (s *TestSuite) SetupSuite() {}

func (s *TestSuite) SetupTest() {}

func (s *TestSuite) TearDownSuite() {}

func (s *TestSuite) TearDownTest() {}

func (s *TestSuite) TestSaveValues() {}

func TestRepoMongodbTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(TestSuite))
}
