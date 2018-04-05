package assign

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Mul node
type Mul struct {
	Variable   node.Node
	Expression node.Node
}

// NewMul node constructor
func NewMul(Variable node.Node, Expression node.Node) *Mul {
	return &Mul{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Mul) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Mul) Walk(v walker.Visitor) {
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
