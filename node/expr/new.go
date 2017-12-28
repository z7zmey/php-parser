package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n New) Name() string {
	return "New"
}

type New struct {
	name      string
	class     node.Node
	arguments []node.Node
}

func NewNew(class node.Node, arguments []node.Node) node.Node {
	return New{
		"New",
		class,
		arguments,
	}
}

func (n New) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.class != nil {
		vv := v.GetChildrenVisitor("class")
		n.class.Walk(vv)
	}

	if n.arguments != nil {
		vv := v.GetChildrenVisitor("arguments")
		for _, nn := range n.arguments {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
