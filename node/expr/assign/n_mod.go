package assign

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Mod node
type Mod struct {
	Variable   node.Node
	Expression node.Node
}

// NewMod node constructor
func NewMod(Variable node.Node, Expression node.Node) *Mod {
	return &Mod{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Mod) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Mod) Walk(v walker.Visitor) {
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
