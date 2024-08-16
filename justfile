test:
  go test -race ./...

clean-cache:
  go clean -testcache