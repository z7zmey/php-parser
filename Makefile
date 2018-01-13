PHPFILE=example.php

all: ./parser/parser.go ./parser/scanner.go
	rm -f y.output
	gofmt -l -s -w *.go
	go build

run: all
	./php-parser $(PHPFILE)

test: all
	go test ./... --cover

./parser/scanner.go: ./parser/scanner.l
	golex -o $@ $<

./parser/parser.go: ./parser/parser.y
	goyacc -o $@ $<
