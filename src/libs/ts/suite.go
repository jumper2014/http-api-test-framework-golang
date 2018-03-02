package ts

import (
	"testing"
	"github.com/stretchr/testify/suite"
)


type myTestSuite struct {
	suite.Suite
}

func (suite *myTestSuite) SetupSuite() {
}

func (suite *myTestSuite) TearDownSuite() {
}

func (suite *myTestSuite) SetupTest() {
}

func (suite *myTestSuite) TearDownTest() {
}

func (suite *myTestSuite) TestMyFunc() {
}

func TestMyTestSuite(t *testing.T) {
	tests := new(myTestSuite)
	suite.Run(t, tests)
}