package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type StaticVar struct {
	attributes map[string]interface{}
	position   *node.Position
	variable   node.Node
	expr       node.Node
}

func NewStaticVar(variable node.Node, expr node.Node) node.Node {
	return &StaticVar{
		map[string]interface{}{},
		nil,
		variable,
		expr,
	}
}

func (n StaticVar) Attributes() map[string]interface{} {
	return n.attributes
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

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
