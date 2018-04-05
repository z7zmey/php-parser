package assign

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Assign node
type Assign struct {
	Variable   node.Node
	Expression node.Node
}

// NewAssign node constructor
func NewAssign(Variable node.Node, Expression node.Node) *Assign {
	return &Assign{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Assign) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Assign) Walk(v walker.Visitor) {
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
