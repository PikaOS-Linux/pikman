.DEFAULT_GOAL := build
build:
	go build -ldflags="-s -w" -o pikman
install:
	mkdir -p $(DESTDIR)/usr/bin/
	install -m 755 pikman $(DESTDIR)/usr/bin/
