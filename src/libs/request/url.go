// Url formatter
// author: Zeng Yue Tian
// date: 2017/12/20

package request

import "strconv"

func CreateUrl(host string, port int, uri string) string  {

	fullUrl := "http://"+host+":"+strconv.Itoa(port)+uri
	return fullUrl
	
}
