.PHONY: build
build:
	go build .

.PHONY: run
run:
	./cli-bnf-token cube.testserver.mail.ru 4995 abracadabra test

test:
	go build .
	go test -run '' -v -cover

.DEFAULT_GOAL := test
