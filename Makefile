PHPFILE=example.php

all: ./php5/php5.go ./php7/php7.go ./scanner/scanner.go
	rm -f y.output
	gofmt -l -s -w *.go
	go build

run: all
	./php-parser $(PHPFILE)

test: all
	go test ./... --cover

./scanner/scanner.go: ./scanner/scanner.l
	golex -o $@ $<

./php5/php5.go: ./php5/php5.y
	goyacc -o $@ $<

./php7/php7.go: ./php7/php7.y
	goyacc -o $@ $<
