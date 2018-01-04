package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Foreach struct {
	position *node.Position
	ByRef    bool
	Expr     node.Node
	Key      node.Node
	Variable node.Node
	Stmt     node.Node
}

func NewForeach(Expr node.Node, Key node.Node, Variable node.Node, Stmt node.Node, ByRef bool) node.Node {
	return &Foreach{
		nil,
		ByRef,
		Expr,
		Key,
		Variable,
		Stmt,
	}
}

func (n Foreach) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ByRef": n.ByRef,
	}
}

func (n Foreach) Position() *node.Position {
	return n.position
}

func (n Foreach) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Foreach) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	if n.Key != nil {
		vv := v.GetChildrenVisitor("Key")
		n.Key.Walk(vv)
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
