package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Pow struct {
	AssignOp
}

func NewPow(variable node.Node, expression node.Node) node.Node {
	return Pow{
		AssignOp{
			"AssignPow",
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n Pow) Name() string {
	return "Pow"
}

func (n Pow) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Pow) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Pow) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Pow) Position() *node.Position {
	return n.position
}

func (n Pow) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Pow) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.expression != nil {
		vv := v.GetChildrenVisitor("expression")
		n.expression.Walk(vv)
	}

	v.LeaveNode(n)
}
