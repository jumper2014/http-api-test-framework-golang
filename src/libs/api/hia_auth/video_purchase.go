// video purchase API
// author: Zeng YueTian
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

func VideoPurchase(videoID string, body VideoPurchaseBody) (*http.Response, error) {
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
type VideoPurchaseBody struct {
	userID     	uint64
	userName	string
	videoName	string
	action   	string
	url   		string

}
