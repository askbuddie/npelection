[ ! -d "./build" ] && mkdir build;
GOOS=linux GOARCH=amd64 go build -o ./build/linux-amd64/npelection ./main.go
GOOS=linux GOARCH=arm go build -o ./build/linux-arm/npelection ./main.go
GOOS=windows GOARCH=amd64 go build -o ./build/windows-amd64/npelection.exe ./main.go
GOOS=windows GOARCH=arm64 go build -o ./build/windows-386/npelection.exe ./main.go
GOOS=darwin GOARCH=amd64 go build -o ./build/darwin-amd64/npelection ./main.go
GOOS=darwin GOARCH=arm64 go build -o ./build/darwin-arm64/npelection ./main.go
