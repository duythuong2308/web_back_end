# run this script in project root directory

# Vet examines Go source code and reports suspicious constructs
go vet ./...

go fmt ./...

# Run all unittests, include some network tests.
# Run `go test` in `pkg/core` for only logic tests
go clean -testcache &&\
    go test -v ./... | grep FAIL -B 5 -A 5
