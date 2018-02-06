package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ClusureUse node
type ClusureUse struct {
	ByRef    bool
	Variable node.Node
}

// NewClusureUse node constuctor
func NewClusureUse(Variable node.Node, ByRef bool) *ClusureUse {
	return &ClusureUse{
		ByRef,
		Variable,
	}
}

// Attributes returns node attributes as map
func (n *ClusureUse) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ByRef": n.ByRef,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ClusureUse) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
