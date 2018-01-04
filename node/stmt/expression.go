package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Expression struct {
	attributes map[string]interface{}
	position   *node.Position
	Expr       node.Node
}

func NewExpression(Expr node.Node) node.Node {
	return &Expression{
		map[string]interface{}{},
		nil,
		Expr,
	}
}

func (n Expression) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Expression) Position() *node.Position {
	return n.position
}

func (n Expression) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Expression) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
