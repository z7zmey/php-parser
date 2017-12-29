package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PreInc struct {
	name     string
	variable node.Node
}

func NewPreInc(variable node.Node) node.Node {
	return PreInc{
		"PreInc",
		variable,
	}
}

func (n PreInc) Name() string {
	return "PreInc"
}

func (n PreInc) Attributes() map[string]interface{} {
	return nil
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
