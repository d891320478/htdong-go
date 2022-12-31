CGO_ENABLED=0 GOOS=linux/windows/darwin GOARCH=amd64/arm64 go build -o Name -x src/Name.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o usr.exe -x src/Main.go