package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// PreInc node
type PreInc struct {
	Variable node.Node
}

// NewPreInc node constructor
func NewPreInc(Variable node.Node) *PreInc {
	return &PreInc{
		Variable,
	}
}

// Attributes returns node attributes as map
func (n *PreInc) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *PreInc) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
