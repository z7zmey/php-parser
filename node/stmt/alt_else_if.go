package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type AltElseIf struct {
	attributes map[string]interface{}
	position   *node.Position
	cond       node.Node
	stmt       node.Node
}

func NewAltElseIf(cond node.Node, stmt node.Node) node.Node {
	return &AltElseIf{
		map[string]interface{}{},
		nil,
		cond,
		stmt,
	}
}

func (n AltElseIf) Attributes() map[string]interface{} {
	return n.attributes
}

func (n AltElseIf) Position() *node.Position {
	return n.position
}

func (n AltElseIf) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n AltElseIf) Walk(v node.Visitor) {
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
