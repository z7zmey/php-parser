package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Static struct {
	Vars []node.Node
}

func NewStatic(Vars []node.Node) *Static {
	return &Static{
		Vars,
	}
}

func (n *Static) Attributes() map[string]interface{} {
	return nil
}

func (n *Static) Walk(v node.Visitor) {
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
