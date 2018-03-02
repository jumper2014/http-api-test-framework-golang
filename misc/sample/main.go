// main program
// author: Zeng YueTian
// create html report

package main

import (
    "net/http"
    "encoding/json"
    "fmt"
    "bytes"
    "io/ioutil"
    "unsafe"
)
 
 
func main() {
    age := make(map[string]interface{})
    age["name"] = "xiaohui"
    bytesData, err := json.Marshal(age)
    if err != nil {
        fmt.Println(err.Error() )
        return
    }
    reader := bytes.NewReader(bytesData)
    url := "http://localhost:8000/age"
    request, err := http.NewRequest("POST", url, reader)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    request.Header.Set("Content-Type", "application/json;charset=UTF-8")
    client := http.Client{}
    resp, err := client.Do(request)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    respBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    //byte数组直接转成string，优化内存
    str := (*string)(unsafe.Pointer(&respBytes))
    fmt.Println(*str)
}


