package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ShortArray struct {
	name  string
	items []node.Node
}

func NewShortArray(items []node.Node) node.Node {
	return ShortArray{
		"ShortArray",
		items,
	}
}

func (n ShortArray) Name() string {
	return "ShortArray"
}

func (n ShortArray) Attributes() map[string]interface{} {
	return nil
}

func (n ShortArray) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.items != nil {
		vv := v.GetChildrenVisitor("items")
		for _, nn := range n.items {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
