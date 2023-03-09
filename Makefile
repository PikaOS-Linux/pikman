.DEFAULT_GOAL := build
build:
	go build -ldflags="-s -w" -o pikman
install:
	install -m 755 pikman $(DESTDIR)/usr/bin/
