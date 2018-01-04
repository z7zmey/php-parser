package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseXor struct {
	AssignOp
}

func NewBitwiseXor(variable node.Node, expression node.Node) node.Node {
	return &BitwiseXor{
		AssignOp{
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n BitwiseXor) Attributes() map[string]interface{} {
	return n.attributes
}

func (n BitwiseXor) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n BitwiseXor) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n BitwiseXor) Position() *node.Position {
	return n.position
}

func (n BitwiseXor) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n BitwiseXor) Walk(v node.Visitor) {
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
