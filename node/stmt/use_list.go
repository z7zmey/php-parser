package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type UseList struct {
	UseType node.Node
	Uses    []node.Node
}

func NewUseList(UseType node.Node, Uses []node.Node) *UseList {
	return &UseList{
		UseType,
		Uses,
	}
}

func (n *UseList) Attributes() map[string]interface{} {
	return nil
}

func (n *UseList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.UseType != nil {
		vv := v.GetChildrenVisitor("UseType")
		n.UseType.Walk(vv)
	}

	if n.Uses != nil {
		vv := v.GetChildrenVisitor("Uses")
		for _, nn := range n.Uses {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
