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

A Parser for PHP written in Go inspired by [Nikic PHP Parser](https://github.com/nikic/PHP-Parser)

## Features:
- Fully support PHP7 syntax (PHP5 in future)
- Abstract syntax tree representation
- Traversing AST

## Example
```Golang
package main

import (
	"bytes"

	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/visitor"
)

func main() {
	src := bytes.NewBufferString(`<? echo "Hello world";`)
	nodes, comments, positions := parser.Parse(src, "example.php")

	dumper := visitor.Dumper{"", comments, positions}
	nodes.Walk(dumper)
}

```

## Roadmap
- [X] Lexer
- [x] PHP 7 syntax analyzer (completely)
- [x] AST nodes
- [x] AST visitor
- [x] AST dumper
- [x] node position
- [x] handling comments
- [ ] Tests
- [ ] Error handling
- [ ] Stabilize api
- [ ] Documentation
- [ ] PHP 5 syntax analyzer
- [ ] Code flow graph
- [ ] Pretty printer
