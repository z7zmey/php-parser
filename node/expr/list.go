package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type List struct {
	name       string
	attributes map[string]interface{}
	items      []node.Node
}

func NewList(items []node.Node) node.Node {
	return List{
		"List",
		map[string]interface{}{},
		items,
	}
}

func (n List) Name() string {
	return "List"
}

func (n List) Attributes() map[string]interface{} {
	return n.attributes
}

func (n List) Walk(v node.Visitor) {
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
