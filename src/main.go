// main program
// author: Zeng YueTian
// create html report for test cases

package main

import (
    "os/exec"
    "fmt"
)
 
 
func main() {

    // install and place your junit2html to PATH
    // https://pypi.python.org/pypi/junit2html
    junit2html := "C:\\Python27\\Scripts\\junit2html"
    fmt.Println(junit2html)

    xml := "F:\\git\\github\\http-api-test-framework-golang\\src\\testsuites\\suite\\report.xml"
    command := "python " + junit2html + " " + xml

    transferCmd := exec.Command("cmd.exe", "/c", command)
    transferOut, err := transferCmd.Output()
    if err != nil {
        panic(err)
    }

    fmt.Println(string(transferOut))


}


