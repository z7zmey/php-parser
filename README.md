PHP Parser written in Go
========================

<img src="./parser.jpg" alt="PHP Parser written in Go" width="980"/>

[![GoDoc](https://godoc.org/github.com/z7zmey/php-parser?status.svg)](https://godoc.org/github.com/z7zmey/php-parser)
[![Build Status](https://travis-ci.org/z7zmey/php-parser.svg?branch=master)](https://travis-ci.org/z7zmey/php-parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/z7zmey/php-parser)](https://goreportcard.com/report/github.com/z7zmey/php-parser)
[![Maintainability](https://api.codeclimate.com/v1/badges/950783b2e739db26e0ed/maintainability)](https://codeclimate.com/github/z7zmey/php-parser/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/950783b2e739db26e0ed/test_coverage)](https://codeclimate.com/github/z7zmey/php-parser/test_coverage)

This project uses [goyacc](https://godoc.org/golang.org/x/tools/cmd/goyacc) and [ragel](https://www.colm.net/open-source/ragel/) tools to create PHP parser. It parses source code into [AST](https://en.wikipedia.org/wiki/Abstract_syntax_tree). It can be used to write static analysis, refactoring, metrics, code style formatting tools.

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
	"log"
	"os"

	"github.com/z7zmey/php-parser/pkg/cfg"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/parser"
	"github.com/z7zmey/php-parser/pkg/version"
	"github.com/z7zmey/php-parser/pkg/visitor/dumper"
)

func main() {
	src := []byte(`<? echo "Hello world";`)

	// Error handler

	var parserErrors []*errors.Error
	errorHandler := func(e *errors.Error) {
		parserErrors = append(parserErrors, e)
	}

	// Parse

	rootNode, err := parser.Parse(src, cfg.Config{
		Version:          &version.Version{Major: 5, Minor: 6},
		ErrorHandlerFunc: errorHandler,
	})

	if err != nil {
		log.Fatal("Error:" + err.Error())
	}

	// Dump

	goDumper := dumper.NewDumper(os.Stdout).
		WithTokens().
		WithPositions()

	rootNode.Accept(goDumper)
}
```

Roadmap
-------

- Control Flow Graph (CFG)
- PHP8

Install
-------

```
go get github.com/z7zmey/php-parser/cmd/php-parser
```

CLI
---

```
php-parser [flags] <path> ...
```

| flag    | type   | description                       |
| ------- | ------ | --------------------------------- |
| -p      | bool   | print filepath                    |
| -e      | bool   | print errors                      |
| -d      | bool   | dump in golang format             |
| -r      | bool   | resolve names                     |
| -prof   | string | start profiler: [cpu, mem, trace] |
| -phpver | string | php version (default: 7.4)        |

Namespace resolver
------------------

Namespace resolver is a visitor that resolves nodes fully qualified name and saves into `map[node.Node]string` structure

- For `Class`, `Interface`, `Trait`, `Function`, `Constant` nodes it saves name with current namespace.
- For `Name`, `Relative`, `FullyQualified` nodes it resolves `use` aliases and saves a fully qualified name.
