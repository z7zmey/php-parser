package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type List struct {
	Items []node.Node
}

func NewList(Items []node.Node) *List {
	return &List{
		Items,
	}
}

func (n *List) Attributes() map[string]interface{} {
	return nil
}

func (n *List) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Items != nil {
		vv := v.GetChildrenVisitor("Items")
		for _, nn := range n.Items {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
