<!--
  Title: PHP Parser
  Description: A Parser for PHP written in Go.
  Author: Slizov Vadim
  Keywords: go golang php php-parser ast
  -->

<img src="./parser.jpg" alt="A parser for PHP written in Go" width="980"/>

[![Go Report Card](https://goreportcard.com/badge/github.com/z7zmey/php-parser)](https://goreportcard.com/report/github.com/z7zmey/php-parser)
[![Exago](https://api.exago.io:443/badge/tests/github.com/z7zmey/php-parser)](https://exago.io/project/github.com/z7zmey/php-parser)
[![Exago](https://api.exago.io:443/badge/cov/github.com/z7zmey/php-parser)](https://exago.io/project/github.com/z7zmey/php-parser)
[![GoDoc](https://godoc.org/github.com/z7zmey/php-parser?status.svg)](https://godoc.org/github.com/z7zmey/php-parser)

#### Try it online: [demo](https://php-parser.com)

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

## CLI dumper

```
$GOPATH/bin/php-parser /path/to/file/or/dir
```

## Namespace resolver

Namespace resolver is a visitor that traverses nodes and resolves nodes fully qualified name.
It does not change AST but collects resolved names into `map[node.Node]string`

- For `Class`, `Interface`, `Trait`, `Function`, `ConstList` nodes collects name with current namespace.
- For `Name`, `Relative`, `FullyQualified` nodes resolves `use` aliases and collects a fully qualified name.

## Pretty printer

```Golang
nodes := &stmt.StmtList{
	Stmts: []node.Node{
		&stmt.Namespace{
			NamespaceName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
		},
		&stmt.Class{
			Modifiers: []node.Node{&node.Identifier{Value: "abstract"}},
			ClassName: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Bar"}}},
			Extends: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Baz"}}},
			Stmts: []node.Node{
				&stmt.ClassMethod{
					Modifiers:  []node.Node{&node.Identifier{Value: "public"}},
					MethodName: &node.Identifier{Value: "greet"},
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
}

file := os.Stdout
p := printer.NewPrinter(file, "    ")
p.PrintFile(nodes)
```

Output:
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
- [x] Pretty printer
- [ ] PhpDocComment parser
- [ ] Error handling
- [ ] Stabilize api
- [ ] Documentation
- [ ] Code flow graph
