buildProxy:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myproxy/proxy myproxy/myproxy.go