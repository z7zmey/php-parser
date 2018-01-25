PHPFILE=example.php

all: ./php7parser/parser.go ./scanner/scanner.go
	rm -f y.output
	gofmt -l -s -w *.go
	go build

run: all
	./php-parser $(PHPFILE)

test: all
	go test ./... --cover

./scanner/scanner.go: ./scanner/scanner.l
	golex -o $@ $<

./php7parser/parser.go: ./php7parser/parser.y
	goyacc -o $@ $<
