package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Default struct {
	Stmts []node.Node
}

func NewDefault(Stmts []node.Node) *Default {
	return &Default{
		Stmts,
	}
}

func (n *Default) Attributes() map[string]interface{} {
	return nil
}

func (n *Default) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
