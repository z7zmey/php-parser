package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Mul struct {
	BinaryOp
}

func NewMul(variable node.Node, expression node.Node) node.Node {
	return Mul{
		BinaryOp{
			"BinaryMul",
			map[string]interface{}{},
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

func (n Mul) Walk(v node.Visitor) {
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
