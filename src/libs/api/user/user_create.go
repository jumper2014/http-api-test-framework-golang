// Account register API
// author: Zeng Yue Tian
// date: 2017/12/22
// done

package user

import (
	"libs/request"
	"libs/constant"
	"net/http"
	"encoding/json"
	"fmt"
	"bytes"
)

func UserCreate(body UserCreateBody) (*http.Response, error) {
	uri := "/user"

	bytesData, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	reader := bytes.NewReader(bytesData)

	return request.SendHttpRequest(
		constant.POST,
		constant.ApiServer,
		constant.Port,
		uri,
		constant.HeaderContentAcceptTypeJson,
		reader)
}

type UserCreateBody struct {
	Name string `json:"name"`
	Age  int `json:"age"`
}
