package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ClosureUse node
type ClosureUse struct {
	ByRef    bool
	Variable node.Node
}

// NewClosureUse node constructor
func NewClosureUse(Variable node.Node, ByRef bool) *ClosureUse {
	return &ClosureUse{
		ByRef,
		Variable,
	}
}

// Attributes returns node attributes as map
func (n *ClosureUse) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ByRef": n.ByRef,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ClosureUse) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
