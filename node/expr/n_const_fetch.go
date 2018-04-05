package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ConstFetch node
type ConstFetch struct {
	Constant node.Node
}

// NewConstFetch node constructor
func NewConstFetch(Constant node.Node) *ConstFetch {
	return &ConstFetch{
		Constant,
	}
}

// Attributes returns node attributes as map
func (n *ConstFetch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ConstFetch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Constant != nil {
		vv := v.GetChildrenVisitor("Constant")
		n.Constant.Walk(vv)
	}

	v.LeaveNode(n)
}
