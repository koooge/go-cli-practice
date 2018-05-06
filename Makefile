build:
	vgo build -o go-cli-practice

test:
	go test ./...

cross:
	GOOS=linux GOARCH=amd64 vgo build -o go-cli-practice_linux_amd64
	GOOS=darwin GOARCH=amd64 vgo build -o go-cli-practice_linux_amd64
	GOOS=windows GOARCH=amd64 vgo build -o go-cli-practice_linux_amd64

.PHONY: build test cross
