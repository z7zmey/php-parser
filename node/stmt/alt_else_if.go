package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type AltElseIf struct {
	name string
	cond node.Node
	stmt node.Node
}

func NewAltElseIf(cond node.Node, stmt node.Node) node.Node {
	return AltElseIf{
		"AltElseIf",
		cond,
		stmt,
	}
}

func (n AltElseIf) Name() string {
	return "AltElseIf"
}

func (n AltElseIf) Attributes() map[string]interface{} {
	return nil
}

func (n AltElseIf) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.cond != nil {
		vv := v.GetChildrenVisitor("cond")
		n.cond.Walk(vv)
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
