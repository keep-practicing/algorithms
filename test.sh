function test() {
    go clean -testcache
    go test ./...
}

test
