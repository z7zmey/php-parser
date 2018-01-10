package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ShortList struct {
	Items []node.Node
}

func NewShortList(Items []node.Node) *ShortList {
	return &ShortList{
		Items,
	}
}

func (n *ShortList) Attributes() map[string]interface{} {
	return nil
}

func (n *ShortList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Items != nil {
		vv := v.GetChildrenVisitor("Items")
		for _, nn := range n.Items {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
