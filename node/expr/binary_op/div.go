package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Div struct {
	BinaryOp
}

func NewDiv(variable node.Node, expression node.Node) node.Node {
	return Div{
		BinaryOp{
			"BinaryDiv",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n Div) Name() string {
	return "Div"
}

func (n Div) Attributes() map[string]interface{} {
	return n.attributes
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
