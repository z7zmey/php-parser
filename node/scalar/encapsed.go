package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type Encapsed struct {
	name  string
	parts []node.Node
}

func NewEncapsed(parts []node.Node) node.Node {
	return Encapsed{
		"Encapsed",
		parts,
	}
}

func (n Encapsed) Name() string {
	return "Encapsed"
}

func (n Encapsed) Attributes() map[string]interface{} {
	return nil
}

func (n Encapsed) Walk(v node.Visitor) {
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
