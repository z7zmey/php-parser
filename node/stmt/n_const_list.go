package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ConstList node
type ConstList struct {
	Consts []node.Node
}

// NewConstList node constructor
func NewConstList(Consts []node.Node) *ConstList {
	return &ConstList{
		Consts,
	}
}

// Attributes returns node attributes as map
func (n *ConstList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ConstList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Consts != nil {
		vv := v.GetChildrenVisitor("Consts")
		for _, nn := range n.Consts {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
