package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ShortArray struct {
	Items []node.Node
}

func NewShortArray(Items []node.Node) *ShortArray {
	return &ShortArray{
		Items,
	}
}

func (n *ShortArray) Attributes() map[string]interface{} {
	return nil
}

func (n *ShortArray) Walk(v node.Visitor) {
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
