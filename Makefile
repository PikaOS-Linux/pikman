.DEFAULT_GOAL := build
build:
	go build -ldflags="-s -w" -o pikman
install:
	install -m 644 pikman /usr/bin/
