package node

import (
	"github.com/z7zmey/php-parser/walker"
)

// ArgumentList node
type ArgumentList struct {
	Arguments []Node
}

// NewArgumentList node constructor
func NewArgumentList(Arguments []Node) *ArgumentList {
	return &ArgumentList{
		Arguments,
	}
}

// Attributes returns node attributes as map
func (n *ArgumentList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ArgumentList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Arguments != nil {
		vv := v.GetChildrenVisitor("Arguments")
		for _, nn := range n.Arguments {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
