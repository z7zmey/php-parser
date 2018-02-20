/*

## Features:

	Fully support PHP5 and PHP7 syntax
	Abstract syntax tree representation
	Traversing AST

## Install

	go get github.com/z7zmey/php-parser

Example

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
*/
package main // import "github.com/z7zmey/php-parser"
