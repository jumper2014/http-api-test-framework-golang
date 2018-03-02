// 注册普通用户
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

func TestUserRegisterEmailOptional(t *testing.T) {

	database.UserDeleteByName(constant.AdminUser)

	var body map[string]interface{}
	body = make(map[string]interface {})
	body["UserID"] = constant.AdminID
	body["UserName"] = constant.AdminUser
	body["Password"] = constant.AdminPassword
	body["UserType"] = constant.UserTypeCommonUser

	resp, err := hia_auth.UserRegister(constant.UserTypeCommonUser, body)

	if err != nil {
		t.Fatal("\t\tRequest send fail", err)
	}
	defer resp.Body.Close()

	expectedStatusCode := constant.StatusCode201
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("\t\tShould receive a \"%d\" status.%v", expectedStatusCode, resp.StatusCode)
	}

	if !database.UserFindByName(constant.AdminUser) {
		t.Errorf("\t\tAdd Fail")
	}

}
