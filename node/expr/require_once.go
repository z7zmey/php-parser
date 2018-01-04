package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type RequireOnce struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewRequireOnce(Expression node.Node) node.Node {
	return &RequireOnce{
		map[string]interface{}{},
		nil,
		Expression,
	}
}

func (n RequireOnce) Attributes() map[string]interface{} {
	return n.attributes
}

func (n RequireOnce) Position() *node.Position {
	return n.position
}

func (n RequireOnce) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n RequireOnce) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
