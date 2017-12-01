# Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# blame: jnml, labs.nic.cz

all: parser.go scanner.go
	rm -f y.output
	gofmt -l -s -w *.go
	go build

run: all
	./php-parser example.php

scanner.go: scanner.l
	golex -o $@ $<

parser.go: parser.y
	goyacc -o $@ $<

clean:
	rm -f php-parser.go lex.yy.go y.output *~

nuke: clean
	rm -f example
