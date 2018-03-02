// main program
// author: Zeng YueTian
// create html report for test cases

package main

import (
    "os/exec"
    "fmt"
)
 
 
func main() {

    junit2html := "C:\\Python27\\Scripts\\junit2html"
    fmt.Println(junit2html)

    xml := "E:\\git\\qa\\gother\\src\\testsuites\\suite\\report.xml"
    command := "python " + junit2html + " " + xml





    transferCmd := exec.Command("cmd.exe", "/c", command)
    transferOut, err := transferCmd.Output()
    if err != nil {
        panic(err)
    }

    fmt.Println(string(transferOut))


}


