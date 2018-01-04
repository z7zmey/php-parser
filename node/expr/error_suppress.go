package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ErrorSuppress struct {
	position *node.Position
	Expr     node.Node
}

func NewErrorSuppress(Expression node.Node) node.Node {
	return &ErrorSuppress{
		nil,
		Expression,
	}
}

func (n ErrorSuppress) Attributes() map[string]interface{} {
	return nil
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

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
