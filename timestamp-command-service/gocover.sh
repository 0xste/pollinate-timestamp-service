go test -race -tags musl -coverprofile cover.out `go list ./... | grep -v cucumber`
go tool cover -func cover.out | grep total | awk '{print "Coverage: " $3}'
go get gotest.tools/gotestsum
gotestsum --debug --junitfile unit-tests.xml `go list ./... | grep -v cucumber` -- -tags=musl