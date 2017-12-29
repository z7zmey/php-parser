package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Concat struct {
	AssignOp
}

func NewConcat(variable node.Node, expression node.Node) node.Node {
	return Concat{
		AssignOp{
			"AssignConcat",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n Concat) Name() string {
	return "Concat"
}

func (n Concat) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Concat) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.expression != nil {
		vv := v.GetChildrenVisitor("expression")
		n.expression.Walk(vv)
	}

	v.LeaveNode(n)
}
