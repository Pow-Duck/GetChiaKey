build_linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build  -o get_chia_key -ldflags "-s -w" cmd/main.go

build_win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o get_chia_key.exe -ldflags "-s -w -H windowsgui -buildmode=exe" cmd/main.go