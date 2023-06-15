go.mod              - project & dependencies
pkgname.go          - package code (use any name instead of pkgname)
doc.go              - package level documentation
pkgname_test.go     - test & benchmark file
example_test.go     - testable example
submod              - folder for submodule (call it any name you want)
- submod.go
- submod_test.go
testdata            - folder for test data
- tokens.json       - example file of test data


-------------------------------

go test -v -race -cover ./...