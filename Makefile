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
	sed -i '' -e 's/yyErrorVerbose = false/yyErrorVerbose = true/g' ./php7/php7.go
	sed -i '' -e 's/yyErrorVerbose = false/yyErrorVerbose = true/g' ./php5/php5.go
	rm -f y.output

./scanner/scanner.go: ./scanner/scanner.l
	golex -o $@ $<

./php5/php5.go: ./php5/php5.y
	goyacc -o $@ $<

./php7/php7.go: ./php7/php7.y
	goyacc -o $@ $<

cpu_pprof:
	GOGC=off go test -cpuprofile cpu.prof -bench=. -benchtime=20s ./php7
	go tool pprof ./php7.test cpu.prof

mem_pprof:
	GOGC=off go test -memprofile mem.prof -bench=. -benchtime=20s -benchmem ./php7
	go tool pprof -alloc_objects ./php7.test mem.prof

cpu_pprof_php5:
	GOGC=off go test -cpuprofile cpu.prof -bench=. -benchtime=20s ./php5
	go tool pprof ./php5.test cpu.prof

mem_pprof_php5:
	GOGC=off go test -memprofile mem.prof -bench=. -benchtime=20s -benchmem ./php5
	go tool pprof -alloc_objects ./php5.test mem.prof
