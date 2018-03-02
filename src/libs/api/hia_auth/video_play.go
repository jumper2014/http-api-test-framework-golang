// video play API
// author: Zeng YueTian
// date: 2017/12/22

package hia_auth

import (
	"libs/request"
	"libs/constant"
	"net/http"
)

func VideoPlay(videoID string, username string, url string) (*http.Response, error) {
	uri := "/videos" + "/" + videoID + "?username=" + username + "&url=" + url

	return request.SendHttpRequest(
		constant.GET,
		constant.HiaServer,
		constant.HiaPort,
		uri,
		constant.HeaderAcceptTypeJson,
		nil)
}

type VideoPlayMessage struct {
	Username	uint64
	Video_name 	string
	Url		 	string
}
