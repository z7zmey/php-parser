package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Div) Name() string {
	return "Div"
}

type Div struct {
	AssignOp
}

func NewDiv(variable node.Node, expression node.Node) node.Node {
	return Div{
		AssignOp{
			"AssignDiv",
			variable,
			expression,
		},
	}
}
