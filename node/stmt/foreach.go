package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Foreach struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
	Key        node.Node
	Variable   node.Node
	Stmt       node.Node
}

func NewForeach(expr node.Node, Key node.Node, Variable node.Node, Stmt node.Node, byRef bool) node.Node {
	return &Foreach{
		map[string]interface{}{
			"byRef": byRef,
		},
		nil,
		expr,
		Key,
		Variable,
		Stmt,
	}
}

func (n Foreach) Attributes() map[string]interface{} {
	return n.attributes
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

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
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
