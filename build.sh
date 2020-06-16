# Build Linux Version
go build -o share-cli_linux
# Build Windows Version
GOOS=windows GOARCH=amd64 go build -o share-cli_windows.exe main.go
# Build macOS Version
GOOS=darwin GOARCH=amd64 go build -o share-cli_macOS main.go