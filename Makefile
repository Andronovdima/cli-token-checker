.PHONY: build
build:
	go build .

.PHONY: run
run:
	go build .
	./cli-token-checker cube.testserver.mail.ru 4995 abracadabra test

test:
	go build .
	go test -run '' -v -cover

.DEFAULT_GOAL := test
