.SILENT:

update: fmt
	go get -u ./...

fmt:
	go mod tidy
	$(GOPATH)/bin/gofumpt -l -w .
	$(GOPATH)/bin/goimports -l -w .

prep.macos:
	go install golang.org/x/tools/cmd/goimports@latest
	go install mvdan.cc/gofumpt@latest
