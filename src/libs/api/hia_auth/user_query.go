package hia_auth

import (
	"libs/request"
	"libs/constant"
	"net/http"
)

func UserQuery(indexType string, username string, indexValue string) (*http.Response, error) {
	uri := "/users" + "/" + indexType + "?username="+username+"&index_value="+indexValue

	return request.SendHttpRequest(
		constant.GET,
		constant.HiaServer,
		constant.HiaPort,
		uri,
		constant.HeaderContentAcceptTypeJson,
		nil)
}


