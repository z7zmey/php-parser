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
			variable,
			expression,
		},
	}
}

func (n SmallerOrEqual) Name() string {
	return "SmallerOrEqual"
}

func (n SmallerOrEqual) Attributes() map[string]interface{} {
	return nil
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
