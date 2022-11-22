default: release

release:
	GOOS=linux GOARCH=amd64 go build -o ./build/linux-amd64/npelection ./main.go
	GOOS=linux GOARCH=arm go build -o ./build/linux-arm/npelection ./main.go
	GOOS=windows GOARCH=amd64 go build -o ./build/windows-amd64/npelection.exe ./main.go
	GOOS=windows GOARCH=arm64 go build -o ./build/windows-386/npelection.exe ./main.go
	GOOS=darwin GOARCH=amd64 go build -o ./build/darwin-amd64/npelection ./main.go

lint:
	golangci-lint run
	go mod tidy -v && git --no-pager diff --quiet go.mod

clean:
	rm -rf build

install-linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

install-fumpt:
	go install mvdan.cc/gofumpt@latest

install-tools: install-linter install-fumpt
