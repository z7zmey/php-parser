package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type RequireOnce struct {
	position *node.Position
	Expr     node.Node
}

func NewRequireOnce(Expression node.Node) *RequireOnce {
	return &RequireOnce{
		nil,
		Expression,
	}
}

func (n *RequireOnce) Attributes() map[string]interface{} {
	return nil
}

func (n *RequireOnce) Position() *node.Position {
	return n.position
}

func (n *RequireOnce) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *RequireOnce) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
