package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Minus struct {
	AssignOp
}

func NewMinus(variable node.Node, expression node.Node) node.Node {
	return Minus{
		AssignOp{
			"AssignMinus",
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n Minus) Name() string {
	return "Minus"
}

func (n Minus) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Minus) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Minus) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Minus) Position() *node.Position {
	return n.position
}

func (n Minus) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Minus) Walk(v node.Visitor) {
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
