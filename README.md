<!--
  Title: PHP Parser
  Description: A Parser for PHP written in Go.
  Author: Slizov Vadim
  Keywords: go golang php php-parser ast
  -->

# PHP-Parser

[![Go Report Card](https://goreportcard.com/badge/github.com/z7zmey/php-parser)](https://goreportcard.com/report/github.com/z7zmey/php-parser)
[![Exago](https://api.exago.io:443/badge/tests/github.com/z7zmey/php-parser)](https://exago.io/project/github.com/z7zmey/php-parser)
[![Exago](https://api.exago.io:443/badge/cov/github.com/z7zmey/php-parser)](https://exago.io/project/github.com/z7zmey/php-parser)
[![GoDoc](https://godoc.org/github.com/z7zmey/php-parser?status.svg)](https://godoc.org/github.com/z7zmey/php-parser)

A Parser for PHP written in Go inspired by [Nikic PHP Parser](https://github.com/nikic/PHP-Parser)

## Features:
- Fully support PHP5 and PHP7 syntax
- Abstract syntax tree representation
- Traversing AST
- Namespace resolver

## Install

```
go get github.com/z7zmey/php-parser
```

## Example
```Golang
package main

import (
	"bytes"

	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/visitor"
)

func main() {
	src := bytes.NewBufferString(`<? echo "Hello world";`)
	nodes, comments, positions := php7.Parse(src, "example.php")

	visitor := visitor.Dumper{
		Indent:    "",
		Comments:  comments,
		Positions: positions,
	}
	nodes.Walk(visitor)
}
```

## CLI dumper

```
$GOPATH/bin/php-parser /path/to/file/or/dir
```

## Roadmap
- [X] Lexer
- [x] PHP 7 syntax analyzer
- [x] AST nodes
- [x] AST visitor
- [x] AST dumper
- [x] node position
- [x] handling comments
- [x] PHP 5 syntax analyzer
- [x] Tests
- [x] Namespace resolver
- [ ] PhpDocComment parser
- [ ] Error handling
- [ ] Stabilize api
- [ ] Documentation
- [ ] Pretty printer
- [ ] Code flow graph
