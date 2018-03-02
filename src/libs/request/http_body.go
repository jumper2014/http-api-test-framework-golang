// lib to create json format http request body
// author: Zeng Yue Tian
// date: 2017/12/20

package request

import (
	"encoding/json"
	"fmt"
	"bytes"
	"io"
)

func MakeJsonBody(body map[string]string ) io.Reader{

	//data := make(map[string]interface{})
	//
	//for key, value := range body {
	//	data[key] = value
	//}

	bytesData, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	reader := bytes.NewReader(bytesData)
	return reader
}
