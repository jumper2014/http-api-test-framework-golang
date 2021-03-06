// Account register API
// author: Zeng Yue Tian
// date: 2017/12/22
// done

package user

import (
	"libs/request"
	"libs/constant"
	"net/http"
	"strconv"
)


func UserDelete(userId int) (*http.Response, error) {
	uri := "/user/" + strconv.Itoa(userId)

	return request.SendHttpRequest(
		constant.DELETE,
		constant.ApiServer,
		constant.Port,
		uri,
		constant.HeaderContentAcceptTypeJson,
		nil)
}


