// Account register test cases
// author: Zeng Yue Tian
// date: 2017/12/25
//

package user_register

import (
	"testing"
	"libs/constant"
	"libs/api/hia_auth"
	"libs/database"
)

var bodies = [] hia_auth.UserRegisterBody{
	{
		UserID:   constant.AdminID,
		UserName: constant.AdminUser,
		Password: constant.AdminPassword,
		UserType: constant.UserTypeCommonUser, // TODO
		Email:    constant.AdminEmail,
	},
	{
		UserID:   constant.AdminID,
		UserName: constant.AdminUser,
		Password: constant.AdminPassword,
		UserType: constant.UserTypeCommonUser, // TODO
		Email:    constant.AdminEmail,
	},
}


var expectedStatusCodes = [] int {
	constant.StatusCode201,
	constant.StatusCode400,
}





func TestUserRegister(t *testing.T) {

	database.UserDeleteByName(constant.AdminUser)


	for index, body := range bodies {

		// TODO
		userType := "user"

		resp, err := hia_auth.UserRegister(userType, body)

		if err != nil {
			t.Fatal("\t\tShould be able to make the Get call", constant.CheckMark, err)
		}
		t.Log("\t\tShould be able to make the Get call", constant.CheckMark)

		defer resp.Body.Close()

		expectedStatusCode := expectedStatusCodes[index]
		if resp.StatusCode == expectedStatusCodes[index] {
			t.Logf("\t\tShould receive a \"%d\" status. %v", expectedStatusCode, constant.CheckMark)
		} else {
			t.Errorf("\t\tShould receive a \"%d\" status. %v %v", expectedStatusCode, constant.BallotX, resp.StatusCode)
		}
	}

}
