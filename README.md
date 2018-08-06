<!--
  Title: PHP Parser
  Description: A Parser for PHP written in Go.
  Author: Slizov Vadym
  Keywords: php parser go golang ast
  -->

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
- Namespace resolver
- Able to parse syntax-invalid PHP files

Roadmap
-------

- Pretty printer
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
php-parser [-php5 -noDump] <path> ...
```

Dump AST to stdout.

Example
-------

```Golang
package main

import (
	"fmt"
	"bytes"
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
		Writer:    os.Stdout,
		Indent:    "",
		Comments:  parser.GetComments(),
		Positions: parser.GetPositions(),
	}

	rootNode := parser.GetRootNode()
	rootNode.Walk(visitor)
}
```

Namespace resolver
------------------

Namespace resolver is a visitor that resolves nodes fully qualified name and saves into `map[node.Node]string` structure

- For `Class`, `Interface`, `Trait`, `Function`, `Constant` nodes it saves name with current namespace.
- For `Name`, `Relative`, `FullyQualified` nodes it resolves `use` aliases and saves a fully qualified name.

Parsing syntax-invalid PHP files
--------------------------------

If we try to parse `$a$b;` then the parser triggers error 'syntax error: unexpected T_VARIABLE'. Token `$b` is unexpected, but parser recovers parsing process and returns `$b;` statement to AST, because it is syntactically correct.

Pretty printer [work in progress]
---------------------------------

```Golang
nodes := &stmt.StmtList{
	Stmts: []node.Node{
		&stmt.Namespace{
			NamespaceName: &name.Name{
				Parts: []node.Node{
					&name.NamePart{Value: "Foo"},
				},
			},
		},
		&stmt.Class{
			Modifiers: []node.Node{
				&node.Identifier{Value: "abstract"},
			},
			ClassName: &name.Name{
				Parts: []node.Node{
					&name.NamePart{Value: "Bar"},
				},
			},
			Extends: &stmt.ClassExtends{
				ClassName: &name.Name{
					Parts: []node.Node{
						&name.NamePart{
							Value: "Baz"
						},
					},
				},
			},
			Stmts: []node.Node{
				&stmt.ClassMethod{
					Modifiers: []node.Node{
						&node.Identifier{Value: "public"},
					},
					MethodName: &node.Identifier{Value: "greet"},
					Stmt: &stmt.StmtList{
						Stmts: []node.Node{
							&stmt.Echo{
								Exprs: []node.Node{
									&scalar.String{Value: "'Hello world'"},
								},
							},
						},
					},
				},
			},
		},
	},
}

file := os.Stdout
p := printer.NewPrinter(file, "    ")
p.Print(nodes)
```

It prints to stdout:

```PHP
<?php
namespace Foo;
abstract class Bar extends Baz
{
	public function greet()
	{
		echo 'Hello world';
	}
}
```
