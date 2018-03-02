// video delete API
// author: Zeng Yue Tian
// date: 2017/12/22
//

package hia_auth

import (
	"libs/request"
	"libs/constant"
	"net/http"
	"encoding/json"
	"fmt"
	"bytes"
)

func VideoDelete(videoID string, body VideoDeleteBody) (*http.Response, error) {
	uri := "/videos" + "/" + videoID

	bytesData, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	reader := bytes.NewReader(bytesData)

	return request.SendHttpRequest(
		constant.DELETE,
		constant.HiaServer,
		constant.HiaPort,
		uri,
		constant.HeaderContentAcceptTypeJson,
		reader)

}

type VideoDeleteBody struct {
	userID 		uint64
	videoName	string
	url         string

}
