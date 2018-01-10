package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Declare struct {
	Consts []node.Node
	Stmt   node.Node
}

func NewDeclare(Consts []node.Node, Stmt node.Node) *Declare {
	return &Declare{
		Consts,
		Stmt,
	}
}

func (n *Declare) Attributes() map[string]interface{} {
	return nil
}

func (n *Declare) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Consts != nil {
		vv := v.GetChildrenVisitor("Consts")
		for _, nn := range n.Consts {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
