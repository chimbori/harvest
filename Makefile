.SILENT:

run.w: fmt
	$(GOPATH)/bin/wgo -verbose go run ./cmd/harvest save

run: fmt
	go run ./cmd/harvest save

update: fmt
	go get -u ./...

fmt:
	go mod tidy
	$(GOPATH)/bin/gofumpt -l -w .
	$(GOPATH)/bin/goimports -l -w .

prep.macos:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/bokwoon95/wgo@latest
	go install mvdan.cc/gofumpt@latest
