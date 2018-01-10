package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type Encapsed struct {
	Parts []node.Node
}

func NewEncapsed(Parts []node.Node) *Encapsed {
	return &Encapsed{
		Parts,
	}
}

func (n *Encapsed) Attributes() map[string]interface{} {
	return nil
}

func (n *Encapsed) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Parts != nil {
		vv := v.GetChildrenVisitor("Parts")
		for _, nn := range n.Parts {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}
}
