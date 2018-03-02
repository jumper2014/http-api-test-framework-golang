package response

import (
	"encoding/json"
)

func DecodeJsonResponse(data []byte) MockAgeMessage {

	var result MockAgeMessage

	err := json.Unmarshal(data, &result)

	if err == nil {
		return result
	} else {
		return result
	}

}


func DecodeResponse(data []byte) EmptyMessage {

	var result EmptyMessage

	err := json.Unmarshal(data, &result)

	if err == nil {
		return result
	} else {
		return result
	}

}


