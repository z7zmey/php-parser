package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseOr struct {
	AssignOp
}

func NewBitwiseOr(variable node.Node, expression node.Node) node.Node {
	return BitwiseOr{
		AssignOp{
			"AssignBitwiseOr",
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n BitwiseOr) Name() string {
	return "BitwiseOr"
}

func (n BitwiseOr) Attributes() map[string]interface{} {
	return n.attributes
}

func (n BitwiseOr) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n BitwiseOr) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n BitwiseOr) Position() *node.Position {
	return n.position
}

func (n BitwiseOr) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n BitwiseOr) Walk(v node.Visitor) {
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
