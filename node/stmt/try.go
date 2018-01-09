package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Try struct {
	Stmts   []node.Node
	Catches []node.Node
	Finally node.Node
}

func NewTry(Stmts []node.Node, Catches []node.Node, Finally node.Node) *Try {
	return &Try{
		Stmts,
		Catches,
		Finally,
	}
}

func (n *Try) Attributes() map[string]interface{} {
	return nil
}

func (n *Try) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	if n.Catches != nil {
		vv := v.GetChildrenVisitor("Catches")
		for _, nn := range n.Catches {
			nn.Walk(vv)
		}
	}

	if n.Finally != nil {
		vv := v.GetChildrenVisitor("Finally")
		n.Finally.Walk(vv)
	}

	v.LeaveNode(n)
}
