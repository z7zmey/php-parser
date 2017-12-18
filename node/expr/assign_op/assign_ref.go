package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type AssignRef struct {
	AssignOp
}

func NewAssignRef(variable node.Node, expression node.Node) node.Node {
	return AssignRef{
		AssignOp{
			node.SimpleNode{Name: "AssignRef", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
