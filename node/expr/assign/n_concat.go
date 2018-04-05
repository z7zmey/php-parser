package assign

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Concat node
type Concat struct {
	Variable   node.Node
	Expression node.Node
}

// NewConcat node constructor
func NewConcat(Variable node.Node, Expression node.Node) *Concat {
	return &Concat{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Concat) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Concat) Walk(v walker.Visitor) {
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
