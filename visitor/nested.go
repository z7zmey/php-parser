// Package visitor contains walker.visitor implementations
package visitor

import (
	"github.com/z7zmey/php-parser/ast"
	"github.com/z7zmey/php-parser/ast/linear"
	"github.com/z7zmey/php-parser/ast/nested"
)

type Nested struct {
	stack []nested.Node
}

func NewNested() *Nested {
	return &Nested{}
}

func (d *Nested) VisitNode(stxtree *linear.AST, n linear.Node, depth int) bool {
	if len(d.stack) <= depth+1 {
		d.stack = append(d.stack, nested.Node{})
	}

	d.stack[depth] = nested.Node{
		Type:     n.Type,
		Flags:    n.Flag,
		Position: stxtree.Positions.Get(n.Pos),
		Children: make(map[ast.EdgeType][]nested.Node),
	}

	if depth > 0 {
		d.stack[depth-1].Children[n.Key] = append(d.stack[depth-1].Children[n.Key], d.stack[depth])
	}

	return true
}

func (d *Nested) Get() nested.Node {
	return d.stack[0]
}
