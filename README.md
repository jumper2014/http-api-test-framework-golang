# HTTP API Test Framework of Golang

### Application under Testing
https://github.com/jumper2014/http-api-aut

### Features
- support HTTP API request
- support MySQL operation to verify result
- support test case run and test suite run
- support test report generation

### Run test one by one
    cd src
    cd testsuites
    cd user
    go test

### Run testsuite (setup, teardown)
    cd src
    cd testsuites
    cd suite
    go test

### Create report for Windows System
    run to generate XML report :
    go test -v 2>&1 |go-junit-report > report.xml

    run to generate HTML report:
    cd src
    go run main.go


### Notes
- Please ignore misc/sample folder