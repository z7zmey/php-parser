PHP Parser written in Go
========================

<img src="./parser.jpg" alt="PHP Parser written in Go" width="980"/>

[![GoDoc](https://godoc.org/github.com/z7zmey/php-parser?status.svg)](https://godoc.org/github.com/z7zmey/php-parser)
[![Build Status](https://travis-ci.org/z7zmey/php-parser.svg?branch=master)](https://travis-ci.org/z7zmey/php-parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/z7zmey/php-parser)](https://goreportcard.com/report/github.com/z7zmey/php-parser)
[![Maintainability](https://api.codeclimate.com/v1/badges/950783b2e739db26e0ed/maintainability)](https://codeclimate.com/github/z7zmey/php-parser/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/950783b2e739db26e0ed/test_coverage)](https://codeclimate.com/github/z7zmey/php-parser/test_coverage)

This project uses [goyacc](https://godoc.org/golang.org/x/tools/cmd/goyacc) and [golex](https://github.com/cznic/golex) libraries to parse PHP sources into [AST](https://en.wikipedia.org/wiki/Abstract_syntax_tree). It can be used to write static analysis, refactoring, metrics, code style formatting tools.

#### Try it online: [demo](https://php-parser.com)

Features:
---------

- Fully support PHP 5 and PHP 7 syntax
- Abstract syntax tree (AST) representation
- Traversing AST
- Resolving namespaced names
- Parsing syntax-invalid PHP files
- Saving and printing free-floating comments and whitespaces

Who Uses
--------

[VKCOM/noverify](https://github.com/VKCOM/noverify) - NoVerify is a pretty fast linter for PHP

[quasilyte/phpgrep](https://github.com/quasilyte/phpgrep) - phpgrep is a tool for syntax-aware PHP code search

Usage example
-------

```Golang
package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/visitor"
)

func main() {
	src := bytes.NewBufferString(`<? echo "Hello world";`)

	parser := php7.NewParser(src, "example.php")
	parser.Parse()

	for _, e := range parser.GetErrors() {
		fmt.Println(e)
	}

	visitor := visitor.Dumper{
		Writer: os.Stdout,
		Indent: "",
	}

	rootNode := parser.GetRootNode()
	rootNode.Walk(&visitor)
}
```

Roadmap
-------

- Control Flow Graph (CFG)
- PhpDocComment parser
- Stabilize api

Install
-------

```
go get github.com/z7zmey/php-parser
```

CLI
---

```
php-parser [flags] <path> ...
```

| flag  | type |                description                   |
|-------|------|----------------------------------------------|
| -d    |string| dump format: [custom, go, json, pretty-json] |
| -r    | bool | resolve names                                |
| -ff   | bool | parse and show free floating strings         |
| -prof |string| start profiler: [cpu, mem, trace]            |
| -php5 | bool | parse as PHP5                                |

Dump AST to stdout.

Namespace resolver
------------------

Namespace resolver is a visitor that resolves nodes fully qualified name and saves into `map[node.Node]string` structure

- For `Class`, `Interface`, `Trait`, `Function`, `Constant` nodes it saves name with current namespace.
- For `Name`, `Relative`, `FullyQualified` nodes it resolves `use` aliases and saves a fully qualified name.
