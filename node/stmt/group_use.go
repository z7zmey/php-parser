package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type GroupUse struct {
	UseType node.Node
	pRefix  node.Node
	UseList []node.Node
}

func NewGroupUse(UseType node.Node, pRefix node.Node, UseList []node.Node) *GroupUse {
	return &GroupUse{
		UseType,
		pRefix,
		UseList,
	}
}

func (n *GroupUse) Attributes() map[string]interface{} {
	return nil
}

func (n *GroupUse) SetUseType(UseType node.Node) node.Node {
	n.UseType = UseType
	return n
}

func (n *GroupUse) Walk(v node.Visitor) {
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
