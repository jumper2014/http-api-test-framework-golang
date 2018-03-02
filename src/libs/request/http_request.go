// core lib to send out http request
// author: Zeng Yue Tian
// date: 2017/12/20

package request

import (

	"fmt"

	"net/http"
	"io"

)


func SendHttpRequest(method string, host string, port int, uri string, header map[string]string, body io.Reader) (*http.Response, error){
	// create full url
	url := CreateUrl(host, port, uri)

	// create request
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// add header for request
	for key, value := range header {
		request.Header.Set(key, value)
	}

	// create http client
	client := http.Client{}

	// send out request
	return client.Do(request)

}


func SendHRequest(method string, url string, reader io.Reader) (*http.Response, error){

	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}

	return client.Do(request)
	//resp, err := client.Do(request)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//respBytes, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	////byte数组直接转成string，优化内存
	//str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println(*str)
}

