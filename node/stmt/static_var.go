package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type StaticVar struct {
	position *node.Position
	Variable node.Node
	Expr     node.Node
}

func NewStaticVar(Variable node.Node, Expr node.Node) node.Node {
	return &StaticVar{
		nil,
		Variable,
		Expr,
	}
}

func (n StaticVar) Attributes() map[string]interface{} {
	return nil
}

func (n StaticVar) Position() *node.Position {
	return n.position
}

func (n StaticVar) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n StaticVar) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
