package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Array struct {
	Items []node.Node
}

func NewArray(Items []node.Node) *Array {
	return &Array{
		Items,
	}
}

func (n *Array) Attributes() map[string]interface{} {
	return nil
}

func (n *Array) Walk(v node.Visitor) {
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
