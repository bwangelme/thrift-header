.PHONY: gen

gen:
	thrift --gen 'go:package_prefix=thrift_header/gen-go/' -strict -r tutorial.thrift

build:
	go build -o bin/thrift_header .

server:
	./bin/thrift_header -server

client:
	./bin/thrift_header