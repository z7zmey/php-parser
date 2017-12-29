package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Minus struct {
	BinaryOp
}

func NewMinus(variable node.Node, expression node.Node) node.Node {
	return Minus{
		BinaryOp{
			"BinaryMinus",
			map[string]interface{}{},
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

func (n Minus) Walk(v node.Visitor) {
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
