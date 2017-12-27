package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Concat) Name() string {
	return "Concat"
}

type Concat struct {
	AssignOp
}

func NewConcat(variable node.Node, expression node.Node) node.Node {
	return Concat{
		AssignOp{
			"AssignConcat",
			variable,
			expression,
		},
	}
}
