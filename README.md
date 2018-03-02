# HTTP API Test Framework of Golang

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