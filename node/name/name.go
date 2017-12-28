package name

import (
	"github.com/z7zmey/php-parser/node"
)

func (n NameNode) Name() string {
	return "Name"
}

type NameNode struct {
	name  string
	parts []node.Node
}

func NewName(parts []node.Node) node.Node {
	return NameNode{
		"Name",
		parts,
	}
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
}
