package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n PreInc) Name() string {
	return "PreInc"
}

type PreInc struct {
	name     string
	variable node.Node
}

func NewPreInc(variableession node.Node) node.Node {
	return PreInc{
		"PreInc",
		variableession,
	}
}

func (n PreInc) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	v.LeaveNode(n)
}
