package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ClassConstList node
type ClassConstList struct {
	Modifiers []node.Node
	Consts    []node.Node
}

// NewClassConstList node constructor
func NewClassConstList(Modifiers []node.Node, Consts []node.Node) *ClassConstList {
	return &ClassConstList{
		Modifiers,
		Consts,
	}
}

// Attributes returns node attributes as map
func (n *ClassConstList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ClassConstList) Walk(v walker.Visitor) {
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
