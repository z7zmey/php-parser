package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Pow struct {
	BinaryOp
}

func NewPow(variable node.Node, expression node.Node) node.Node {
	return Pow{
		BinaryOp{
			"BinaryPow",
			variable,
			expression,
		},
	}
}

func (n Pow) Name() string {
	return "Pow"
}

func (n Pow) Attributes() map[string]interface{} {
	return nil
}

func (n Pow) Walk(v node.Visitor) {
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
