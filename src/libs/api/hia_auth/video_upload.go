// video upload API
// author: Zeng Yue Tian
// date: 2017/12/22
// need to be done

package hia_auth

import (
	"libs/request"
	"libs/constant"
	"net/http"
	"encoding/json"
	"fmt"
	"bytes"
)

func VideoUpload(videoID string, body VideoUploadBody) (*http.Response, error) {
	uri := "/videos" + "/" + string(videoID)

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

type VideoUploadBody struct {
	userID		uint64
	videoName 	string
	url 		string
}