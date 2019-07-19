PHPFILE=example.php

all: compile fmt build run

fmt:
	find . -type f -iregex '.*\.go' -exec gofmt -l -s -w '{}' +

build:
	go generate ./...
	go build ./cmd/...

install:
	go generate ./...
	go install ./cmd/...

test:
	go test ./...

cover:
	go test ./... --cover

bench:
	go test -benchmem -bench=. ./parser/php5
	go test -benchmem -bench=. ./parser/php7

compile: ./internal/parser/php5/php5.go ./internal/parser/php7/php7.go ./internal/scanner/scanner.go fmt
	sed -i '' -e 's/yyErrorVerbose = false/yyErrorVerbose = true/g' ./internal/parser/php7/php7.go
	sed -i '' -e 's/yyErrorVerbose = false/yyErrorVerbose = true/g' ./internal/parser/php5/php5.go
	sed -i '' -e 's/\/\/line/\/\/ line/g' ./internal/parser/php7/php7.go
	sed -i '' -e 's/\/\/line/\/\/ line/g' ./internal/parser/php5/php5.go
	rm -f y.output

./internal/scanner/scanner.go: ./internal/scanner/scanner.rl
	ragel -Z -G2 -o $@ $<

./internal/parser/php5/php5.go: ./internal/parser/php5/php5.y
	goyacc -o $@ $<

./internal/parser/php7/php7.go: ./internal/parser/php7/php7.y
	goyacc -o $@ $<

cpu_pprof:
	go test -cpuprofile cpu.pprof -run=^$$ -bench=^BenchmarkPhp7$$ -benchtime=20s ./internal/parser/php7
	go tool pprof ./php7.test cpu.pprof

mem_pprof:
	go test -memprofile mem.pprof  -run=^$$ -bench=^BenchmarkPhp7$$ -benchtime=20s -benchmem ./internal/parser/php7
	go tool pprof -alloc_objects ./php7.test mem.pprof

cpu_pprof_php5:
	go test -cpuprofile cpu.prof -bench=. -benchtime=20s ./internal/parser/php5
	go tool pprof ./php5.test cpu.prof

mem_pprof_php5:
	go test -memprofile mem.prof -bench=. -benchtime=20s -benchmem ./internal/parser/php5
	go tool pprof -alloc_objects ./php5.test mem.prof
