package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type GreaterOrEqual struct {
	BinaryOp
}

func NewGreaterOrEqual(variable node.Node, expression node.Node) node.Node {
	return GreaterOrEqual{
		BinaryOp{
			"BinaryGreaterOrEqual",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n GreaterOrEqual) Name() string {
	return "GreaterOrEqual"
}

func (n GreaterOrEqual) Attributes() map[string]interface{} {
	return n.attributes
}

func (n GreaterOrEqual) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n GreaterOrEqual) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n GreaterOrEqual) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.left != nil {
		vv := v.GetChildrenVisitor("left")
		n.left.Walk(vv)
	}

	if n.right != nil {
		vv := v.GetChildrenVisitor("right")
		n.right.Walk(vv)
	}

	v.LeaveNode(n)
}
