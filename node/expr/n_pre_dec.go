package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// PreDec node
type PreDec struct {
	Variable node.Node
}

// NewPreDec node constructor
func NewPreDec(Variable node.Node) *PreDec {
	return &PreDec{
		Variable,
	}
}

// Attributes returns node attributes as map
func (n *PreDec) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *PreDec) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
