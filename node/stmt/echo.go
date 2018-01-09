package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Echo struct {
	Exprs []node.Node
}

func NewEcho(Exprs []node.Node) *Echo {
	return &Echo{
		Exprs,
	}
}

func (n *Echo) Attributes() map[string]interface{} {
	return nil
}

func (n *Echo) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Exprs != nil {
		vv := v.GetChildrenVisitor("Exprs")
		for _, nn := range n.Exprs {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
