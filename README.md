# testing framework for block chain project

### How to run test
    cd src
    cd testsuites
    cd <suite_dir>
    go test

    or run:
    go test -v 2>&1 |go-junit-report > report.xml

    then run below command to generate HTML report:
    main.go