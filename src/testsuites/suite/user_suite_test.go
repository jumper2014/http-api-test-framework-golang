package user

// Basic imports
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"libs/constant"
	"libs/database"
	"libs/api/user"
	"fmt"

)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type UserTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *UserTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

// All methods that begin with "Test" are run as tests within a
// suite.
//func (suite *MocoTestSuite) TestExample() {
//	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
//}

func (suite *UserTestSuite) TestUserCreate() {
	database.UserDeleteByName(constant.AdminUser)
	var body user.UserCreateBody
	body.Name = constant.AdminUser
	body.Age = constant.AdminAge

	resp, _ := user.UserCreate(body)

	assert.Equal(suite.T(), constant.StatusCode200, resp.StatusCode)
}

// Data driven Testing Sample
func (suite *UserTestSuite) TestUserDDT() {
	fmt.Println("-----------\nDDT\n-----------\n")

	bodys := [] user.UserCreateBody{
		{constant.AdminUser, constant.AdminAge},
		{constant.RootUser, constant.RootAge},
	}

	for _, body := range bodys {
		resp, _ := user.UserCreate(body)
		assert.Equal(suite.T(), constant.StatusCode200, resp.StatusCode)
	}
}


		// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
