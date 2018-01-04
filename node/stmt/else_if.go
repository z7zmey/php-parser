package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ElseIf struct {
	attributes map[string]interface{}
	position   *node.Position
	cond       node.Node
	stmt       node.Node
}

func NewElseIf(cond node.Node, stmt node.Node) node.Node {
	return &ElseIf{
		map[string]interface{}{},
		nil,
		cond,
		stmt,
	}
}

func (n ElseIf) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ElseIf) Position() *node.Position {
	return n.position
}

func (n ElseIf) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ElseIf) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.cond != nil {
		vv := v.GetChildrenVisitor("cond")
		n.cond.Walk(vv)
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
