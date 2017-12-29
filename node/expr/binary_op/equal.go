package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Equal struct {
	BinaryOp
}

func NewEqual(variable node.Node, expression node.Node) node.Node {
	return Equal{
		BinaryOp{
			"BinaryEqual",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n Equal) Name() string {
	return "Equal"
}

func (n Equal) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Equal) Walk(v node.Visitor) {
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
