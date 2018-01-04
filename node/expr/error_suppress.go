package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ErrorSuppress struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewErrorSuppress(expression node.Node) node.Node {
	return &ErrorSuppress{
		map[string]interface{}{},
		nil,
		expression,
	}
}

func (n ErrorSuppress) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ErrorSuppress) Position() *node.Position {
	return n.position
}

func (n ErrorSuppress) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ErrorSuppress) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
