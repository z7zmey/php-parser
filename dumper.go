package main

import (
	"fmt"
	"reflect"

	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type dumper struct {
	indent    string
	comments  comment.Comments
	positions position.Positions
}

func (d dumper) EnterNode(n node.Node) bool {

	fmt.Printf("%v%v", d.indent, reflect.TypeOf(n))
	if p := d.positions[n]; p != nil {
		fmt.Printf(" %v", *p)
	}
	if a := n.Attributes(); len(a) > 0 {
		fmt.Printf(" %v", a)
	}
	fmt.Println()

	if c := d.comments[n]; len(c) > 0 {
		fmt.Printf("%vcomments:\n", d.indent+"  ")
		for _, cc := range c {
			fmt.Printf("%v%q\n", d.indent+"    ", cc)
		}
	}

	return true
}

func (d dumper) GetChildrenVisitor(key string) node.Visitor {
	fmt.Printf("%v%q:\n", d.indent+"  ", key)
	return dumper{d.indent + "    ", d.comments, d.positions}
}

func (d dumper) LeaveNode(n node.Node) {
	// do nothing
}
