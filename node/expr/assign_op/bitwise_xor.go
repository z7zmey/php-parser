package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseXor struct {
	AssignOp
}

func NewBitwiseXor(variable node.Node, expression node.Node) node.Node {
	return BitwiseXor{
		AssignOp{
			"AssignBitwiseXor",
			variable,
			expression,
		},
	}
}

func (n BitwiseXor) Name() string {
	return "BitwiseXor"
}

func (n BitwiseXor) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.Children("variable")
		n.variable.Walk(vv)
	}

	if n.expression != nil {
		vv := v.Children("expression")
		n.expression.Walk(vv)
	}
}
