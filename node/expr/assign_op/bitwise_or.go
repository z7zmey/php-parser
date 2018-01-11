package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

// BitwiseOr node
type BitwiseOr struct {
	AssignOp
}

// NewBitwiseOr node constuctor
func NewBitwiseOr(Variable node.Node, Expression node.Node) *BitwiseOr {
	return &BitwiseOr{
		AssignOp{
			Variable,
			Expression,
		},
	}
}

// Attributes returns node attributes as map
func (n *BitwiseOr) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *BitwiseOr) Walk(v node.Visitor) {
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
