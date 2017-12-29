package name

import (
	"github.com/z7zmey/php-parser/node"
)

type NameNode struct {
	name       string
	attributes map[string]interface{}
	parts      []node.Node
}

func NewName(parts []node.Node) node.Node {
	return NameNode{
		"Name",
		map[string]interface{}{},
		parts,
	}
}

func (n NameNode) Name() string {
	return "Name"
}

func (n NameNode) Attributes() map[string]interface{} {
	return n.attributes
}

func (n NameNode) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.parts != nil {
		vv := v.GetChildrenVisitor("parts")
		for _, nn := range n.parts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
