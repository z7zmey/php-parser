package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Div) Name() string {
	return "Div"
}

type Div struct {
	BinaryOp
}

func NewDiv(variable node.Node, expression node.Node) node.Node {
	return Div{
		BinaryOp{
			"BinaryDiv",
			variable,
			expression,
		},
	}
}

func (n Div) Walk(v node.Visitor) {
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
