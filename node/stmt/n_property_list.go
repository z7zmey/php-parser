package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// PropertyList node
type PropertyList struct {
	Modifiers  []node.Node
	Properties []node.Node
}

// NewPropertyList node constructor
func NewPropertyList(Modifiers []node.Node, Properties []node.Node) *PropertyList {
	return &PropertyList{
		Modifiers,
		Properties,
	}
}

// Attributes returns node attributes as map
func (n *PropertyList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *PropertyList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Modifiers != nil {
		vv := v.GetChildrenVisitor("Modifiers")
		for _, nn := range n.Modifiers {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Properties != nil {
		vv := v.GetChildrenVisitor("Properties")
		for _, nn := range n.Properties {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
