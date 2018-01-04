package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Include struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewInclude(expression node.Node) node.Node {
	return &Include{
		map[string]interface{}{},
		nil,
		expression,
	}
}

func (n Include) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Include) Position() *node.Position {
	return n.position
}

func (n Include) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Include) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
