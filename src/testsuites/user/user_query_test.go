// Create user test cases
// author: Zeng Yue Tian
// date: 2017/12/25
//
package user

import (
	"testing"
	"libs/api/user"
	"libs/constant"

)

func TestUserQuery(t *testing.T) {


	resp, err := user.UserQuery(5)

	if err != nil {
		t.Fatal("\t\tShould be able to make the Get call", constant.CheckMark, err)
	}
	t.Log("\t\tShould be able to make the Get call", constant.CheckMark)

	defer resp.Body.Close()

	expectedStatusCode := constant.StatusCode200
	if resp.StatusCode == expectedStatusCode {
		t.Logf("\t\tShould receive a \"%d\" status. %v", expectedStatusCode, constant.CheckMark)
	} else {
		t.Errorf("\t\tShould receive a \"%d\" status. %v %v", expectedStatusCode, constant.BallotX, resp.StatusCode)
	}

}
