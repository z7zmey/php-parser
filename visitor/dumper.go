// Package visitor contains walker.visitor implementations
package visitor

import (
	"io"
	"strings"
	"text/template"

	"github.com/z7zmey/php-parser/syntaxtree/linkedtree"
)

type tplData struct {
	Indent string
	Node   linkedtree.Node
}

type stackItem struct {
	parent     linkedtree.NodeID
	nextNodeID linkedtree.NodeID
	depth      int
}

type Dumper struct {
	Writer    io.Writer
	edgeStack []linkedtree.EdgeType
}

func (d *Dumper) VisitNode(n linkedtree.Node, depth int) bool {

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
