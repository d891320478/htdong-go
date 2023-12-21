CGO_ENABLED=0 GOOS=linux/windows/darwin GOARCH=amd64/arm64 go build -o build/Name -x src/Name.go

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/startBiliHttp -x src/Main.go