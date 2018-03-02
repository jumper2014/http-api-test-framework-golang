// Account register API
// author: Zeng Yue Tian
// date: 2017/12/22
// done

package hia_auth

import (
	"libs/request"
	"libs/constant"
	"net/http"
	"encoding/json"
	"fmt"
	"bytes"
)

//func UserRegister(userType string, body UserRegisterBody) (*http.Response, error) {
//	uri := "/users" + "/" + userType
//
//	bytesData, err := json.Marshal(body)
//	if err != nil {
//		fmt.Println(err.Error())
//		return nil, err
//	}
//	reader := bytes.NewReader(bytesData)
//
//	return request.SendHttpRequest(
//		constant.POST,
//		constant.HiaServer,
//		constant.HiaPort,
//		uri,
//		constant.HeaderContentAcceptTypeJson,
//		reader)
//}

func UserRegister(userType string, body map[string]interface{}) (*http.Response, error) {
	uri := "/users" + "/" + userType

	bytesData, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	reader := bytes.NewReader(bytesData)

	return request.SendHttpRequest(
		constant.POST,
		constant.HiaServer,
		constant.HiaPort,
		uri,
		constant.HeaderContentAcceptTypeJson,
		reader)
}


type UserRegisterBody struct {
	UserID 		uint64
	UserName 	string
	Password 	string
	UserType 	string
	Email		string
}


