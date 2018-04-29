package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// CaseList node
type CaseList struct {
	Cases []node.Node
}

// NewCaseList node constructor
func NewCaseList(Cases []node.Node) *CaseList {
	return &CaseList{
		Cases,
	}
}

// Attributes returns node attributes as map
func (n *CaseList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *CaseList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cases != nil {
		vv := v.GetChildrenVisitor("Cases")
		for _, nn := range n.Cases {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
