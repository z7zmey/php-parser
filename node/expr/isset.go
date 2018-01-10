package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Isset struct {
	Variables []node.Node
}

func NewIsset(Variables []node.Node) *Isset {
	return &Isset{
		Variables,
	}
}

func (n *Isset) Attributes() map[string]interface{} {
	return nil
}

func (n *Isset) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variables != nil {
		vv := v.GetChildrenVisitor("Variables")
		for _, nn := range n.Variables {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
