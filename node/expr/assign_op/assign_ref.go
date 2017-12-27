package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n AssignRef) Name() string {
	return "AssignRef"
}

type AssignRef struct {
	AssignOp
}

func NewAssignRef(variable node.Node, expression node.Node) node.Node {
	return AssignRef{
		AssignOp{
			"AssignRef",
			variable,
			expression,
		},
	}
}
