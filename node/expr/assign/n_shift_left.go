package assign

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ShiftLeft node
type ShiftLeft struct {
	Variable   node.Node
	Expression node.Node
}

// NewShiftLeft node constructor
func NewShiftLeft(Variable node.Node, Expression node.Node) *ShiftLeft {
	return &ShiftLeft{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *ShiftLeft) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ShiftLeft) Walk(v walker.Visitor) {
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
