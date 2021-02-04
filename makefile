all:
	env GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o go-datagov .