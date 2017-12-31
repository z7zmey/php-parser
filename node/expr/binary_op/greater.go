package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Greater struct {
	BinaryOp
}

func NewGreater(variable node.Node, expression node.Node) node.Node {
	return Greater{
		BinaryOp{
			"BinaryGreater",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n Greater) Name() string {
	return "Greater"
}

func (n Greater) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Greater) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Greater) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Greater) Walk(v node.Visitor) {
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
