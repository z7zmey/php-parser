package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ShortList struct {
	name  string
	items []node.Node
}

func NewShortList(items []node.Node) node.Node {
	return ShortList{
		"ShortList",
		items,
	}
}

func (n ShortList) Name() string {
	return "ShortList"
}

func (n ShortList) Attributes() map[string]interface{} {
	return nil
}

func (n ShortList) Walk(v node.Visitor) {
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
