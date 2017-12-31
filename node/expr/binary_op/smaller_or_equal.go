package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type SmallerOrEqual struct {
	BinaryOp
}

func NewSmallerOrEqual(variable node.Node, expression node.Node) node.Node {
	return SmallerOrEqual{
		BinaryOp{
			"BinarySmallerOrEqual",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n SmallerOrEqual) Name() string {
	return "SmallerOrEqual"
}

func (n SmallerOrEqual) Attributes() map[string]interface{} {
	return n.attributes
}

func (n SmallerOrEqual) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n SmallerOrEqual) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n SmallerOrEqual) Walk(v node.Visitor) {
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
