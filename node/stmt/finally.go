package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Finally struct {
	Stmts []node.Node
}

func NewFinally(Stmts []node.Node) *Finally {
	return &Finally{
		Stmts,
	}
}

func (n *Finally) Attributes() map[string]interface{} {
	return nil
}

func (n *Finally) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
