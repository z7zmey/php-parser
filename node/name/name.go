package name

import (
	"github.com/z7zmey/php-parser/node"
)

type Name struct {
	Parts []node.Node
}

func NewName(Parts []node.Node) *Name {
	return &Name{
		Parts,
	}
}

func (n *Name) Attributes() map[string]interface{} {
	return nil
}

func (n *Name) Walk(v node.Visitor) {
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

	v.LeaveNode(n)
}
