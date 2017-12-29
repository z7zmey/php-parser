package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type LogicalOr struct {
	BinaryOp
}

func NewLogicalOr(variable node.Node, expression node.Node) node.Node {
	return LogicalOr{
		BinaryOp{
			"BinaryLogicalOr",
			variable,
			expression,
		},
	}
}

func (n LogicalOr) Name() string {
	return "LogicalOr"
}

func (n LogicalOr) Attributes() map[string]interface{} {
	return nil
}

func (n LogicalOr) Walk(v node.Visitor) {
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
