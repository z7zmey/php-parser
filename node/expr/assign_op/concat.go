package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

// Concat node
type Concat struct {
	AssignOp
}

// NewConcat node constuctor
func NewConcat(Variable node.Node, Expression node.Node) *Concat {
	return &Concat{
		AssignOp{
			Variable,
			Expression,
		},
	}
}

// Attributes returns node attributes as map
func (n *Concat) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Concat) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Expression != nil {
		vv := v.GetChildrenVisitor("Expression")
		n.Expression.Walk(vv)
	}

	v.LeaveNode(n)
}
