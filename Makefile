.DEFAULT_GOAL := build
build:
	go build -o pikman
install:
	install -m 644 pikman /usr/bin/
