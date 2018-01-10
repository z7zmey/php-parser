package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Unset struct {
	Vars []node.Node
}

func NewUnset(Vars []node.Node) *Unset {
	return &Unset{
		Vars,
	}
}

func (n *Unset) Attributes() map[string]interface{} {
	return nil
}

func (n *Unset) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Vars != nil {
		vv := v.GetChildrenVisitor("Vars")
		for _, nn := range n.Vars {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
