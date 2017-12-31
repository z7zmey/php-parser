package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Mul struct {
	AssignOp
}

func NewMul(variable node.Node, expression node.Node) node.Node {
	return Mul{
		AssignOp{
			"AssignMul",
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n Mul) Name() string {
	return "Mul"
}

func (n Mul) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Mul) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Mul) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Mul) Position() *node.Position {
	return n.position
}

func (n Mul) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Mul) Walk(v node.Visitor) {
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
