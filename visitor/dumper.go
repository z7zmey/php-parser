// Package visitor contains walker.visitor implementations
package visitor

import (
	"io"
	"strings"
	"text/template"

	"github.com/z7zmey/php-parser/ast"
	"github.com/z7zmey/php-parser/ast/linear"
)

type tplData struct {
	Indent string
	Node   linear.Node
}

type stackItem struct {
	parent     linear.NodeID
	nextNodeID linear.NodeID
	depth      int
}

type Dumper struct {
	Writer    io.Writer
	edgeStack []ast.EdgeType
}

func (d *Dumper) VisitNode(n linear.Node, depth int) bool {

	// print key

	if len(d.edgeStack) <= depth {
		d.edgeStack = append(d.edgeStack, 0)
	}

	if d.edgeStack[depth] != n.Key {
		println(strings.Repeat("    ", depth*2-1) + n.Key.String())
	}
	d.edgeStack[depth] = n.Key

	// print node

	tpl, err := template.New("dump").Parse("{{.Indent}}[*{{.Node.Type}}]\n")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(d.Writer, tplData{
		Indent: strings.Repeat("    ", depth*2),
		Node:   n,
	})
	if err != nil {
		panic(err)
	}

	// continue traversing

	return true
}
