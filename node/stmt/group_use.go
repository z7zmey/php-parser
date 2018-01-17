package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// GroupUse node
type GroupUse struct {
	UseType node.Node
	pRefix  node.Node
	UseList []node.Node
}

// NewGroupUse node constuctor
func NewGroupUse(UseType node.Node, pRefix node.Node, UseList []node.Node) *GroupUse {
	return &GroupUse{
		UseType,
		pRefix,
		UseList,
	}
}

// Attributes returns node attributes as map
func (n *GroupUse) Attributes() map[string]interface{} {
	return nil
}

func (n *GroupUse) SetUseType(UseType node.Node) node.Node {
	n.UseType = UseType
	return n
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *GroupUse) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.UseType != nil {
		vv := v.GetChildrenVisitor("UseType")
		n.UseType.Walk(vv)
	}

	if n.pRefix != nil {
		vv := v.GetChildrenVisitor("pRefix")
		n.pRefix.Walk(vv)
	}

	if n.UseList != nil {
		vv := v.GetChildrenVisitor("UseList")
		for _, nn := range n.UseList {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
