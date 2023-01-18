.DEFAULT_GOAL := build
build:
	GOARCH=amd64 GOOS=linux go build -o pikman
install:
	install -m 644 pikman /usr/bin/
