package suite

// Basic imports
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"libs/api/moco"
	"io/ioutil"
	"libs/response"
	"libs/constant"
	"fmt"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type MocoTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *MocoTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

// All methods that begin with "Test" are run as tests within a
// suite.
//func (suite *MocoTestSuite) TestExample() {
//	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
//}

func (suite *MocoTestSuite) TestMocoAge() {
	fmt.Println("-----------\nTestMocoAge\n-----------\n")

	body := map[string]string{
		"name": "xiaohui",
	}

	resp, err := moco.UserAge(body)

	if err != nil {
		fmt.Println("\t\tShould be able to make the Get call",
			constant.CheckMark, err)

	}
	fmt.Println("\t\tShould be able to make the Get call", constant.CheckMark)

	defer resp.Body.Close()

	if resp.StatusCode == constant.StatusCode200 {
		fmt.Printf("\t\tShould receive a \"%d\" status. %v\n", constant.StatusCode200, constant.CheckMark)
	} else {
		fmt.Printf("\t\tShould receive a \"%d\" status. %v %v\n", constant.StatusCode200, constant.BallotX, resp.StatusCode)
	}

	fmt.Println("\t\tHello")
	data, err := ioutil.ReadAll(resp.Body)
	dataStruct := response.DecodeResponse(data)
	fmt.Println(data)
	fmt.Println(dataStruct)
	fmt.Println(dataStruct == response.EmptyJson)

	assert.Equal(suite.T(), "hello", "")
}


func (suite *MocoTestSuite) TestMocoDDT() {
	fmt.Println("-----------\nTestMocoDDT\n-----------\n")

	bodys := [] map[string]string{
		{
			"name": "xiaohui",
		},
		{
			"name": "xiaohui",
		},
	}

	for _, body := range bodys {

		resp, err := moco.UserAge(body)

		if err != nil {
			fmt.Println("\t\tShould be able to make the Get call",
				constant.CheckMark, err)

		}
		fmt.Println("\t\tShould be able to make the Get call", constant.CheckMark)

		defer resp.Body.Close()

		if resp.StatusCode == constant.StatusCode200 {
			fmt.Printf("\t\tShould receive a \"%d\" status. %v\n", constant.StatusCode200, constant.CheckMark)
		} else {
			fmt.Printf("\t\tShould receive a \"%d\" status. %v %v\n", constant.StatusCode200, constant.BallotX, resp.StatusCode)
		}

		fmt.Println("\t\tHello")
		data, err := ioutil.ReadAll(resp.Body)
		dataStruct := response.DecodeJsonResponse(data)

		fmt.Println(dataStruct.Age)
		fmt.Println(dataStruct.Xu)

		expectedMessage := response.MockAgeMessage{
			Age: 24,
			Xu:  true,
		}

		if dataStruct.Age != expectedMessage.Age ||
			dataStruct.Xu != expectedMessage.Xu {
			fmt.Printf("expect age is: %d\n", expectedMessage.Age)
			fmt.Printf("expect xu is: %t\n", expectedMessage.Xu)
			fmt.Printf("expect value is: %#v\n", expectedMessage)
			fmt.Printf("expect value is: %+v\n", expectedMessage)
		}
	}
}


		// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(MocoTestSuite))
}
