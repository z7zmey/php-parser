package main

import (
	"fmt"

	"github.com/z7zmey/php-parser/node"
)

type dumper struct {
	indent string
}

func (d dumper) EnterNode(n node.Node) bool {
	fmt.Printf("%v[%v]:\n", d.indent, n.Name())

	return true
}

func (d dumper) GetChildrenVisitor(key string) node.Visitor {
	fmt.Printf("%v%v:\n", d.indent+". ", key)
	return dumper{d.indent + ". . "}
}

func (d dumper) Scalar(key string, value interface{}) {
	fmt.Printf("%v%v: %v\n", d.indent+". ", key, value)
}

func (d dumper) LeaveNode(n node.Node) {
	// do nothing
}
