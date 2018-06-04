PHPFILE=example.php

all: compile fmt build run

fmt:
	find . -type f -iregex '.*\.go' -exec gofmt -l -s -w '{}' +

build:
	go build

run:
	./php-parser $(PHPFILE)

test:
	go test ./...

cover:
	go test ./... --cover

bench:
	go test -benchmem -bench=. ./php5
	go test -benchmem -bench=. ./php7

compile: ./php5/php5.go ./php7/php7.go ./scanner/scanner.go
	rm -f y.output

./scanner/scanner.go: ./scanner/scanner.l
	golex -o $@ $<

./php5/php5.go: ./php5/php5.y
	goyacc -o $@ $<

./php7/php7.go: ./php7/php7.y
	goyacc -o $@ $<

profile5:
	GOGC=off go test -cpuprofile cpu.prof -memprofile mem.prof -bench=. -benchmem -benchtime=20s ./php5

profile7:
	GOGC=off go test -cpuprofile cpu.prof -memprofile mem.prof -bench=. -benchmem -benchtime=20s ./php7

cpu_pprof:
	go tool pprof cpu.prof

mem_pprof:
	go tool pprof mem.prof
