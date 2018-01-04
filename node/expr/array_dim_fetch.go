package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ArrayDimFetch struct {
	attributes map[string]interface{}
	position   *node.Position
	Variable   node.Node
	Dim        node.Node
}

func NewArrayDimFetch(Variable node.Node, Dim node.Node) node.Node {
	return &ArrayDimFetch{
		map[string]interface{}{},
		nil,
		Variable,
		Dim,
	}
}

func (n ArrayDimFetch) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ArrayDimFetch) Position() *node.Position {
	return n.position
}

func (n ArrayDimFetch) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ArrayDimFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Dim != nil {
		vv := v.GetChildrenVisitor("Dim")
		n.Dim.Walk(vv)
	}

	v.LeaveNode(n)
}
