package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type AltElse struct {
	Stmt node.Node
}

func NewAltElse(Stmt node.Node) *AltElse {
	return &AltElse{
		Stmt,
	}
}

func (n *AltElse) Attributes() map[string]interface{} {
	return nil
}

func (n *AltElse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
