package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Else struct {
	Stmt node.Node
}

func NewElse(Stmt node.Node) *Else {
	return &Else{
		Stmt,
	}
}

func (n *Else) Attributes() map[string]interface{} {
	return nil
}

func (n *Else) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
