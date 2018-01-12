# Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# blame: jnml, labs.nic.cz
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
