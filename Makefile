PHPFILE=example.php

all: compile run

build: 
	gofmt -l -s -w **/*.go
	go build

run: build
	./php-parser $(PHPFILE)

test: all
	go test ./... --cover

compile: ./php5/php5.go ./php7/php7.go ./scanner/scanner.go
	rm -f y.output

./scanner/scanner.go: ./scanner/scanner.l
	golex -o $@ $<

./php5/php5.go: ./php5/php5.y
	goyacc -o $@ $<

./php7/php7.go: ./php7/php7.y
	goyacc -o $@ $<
