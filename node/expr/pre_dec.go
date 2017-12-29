package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PreDec struct {
	name     string
	variable node.Node
}

func NewPreDec(variable node.Node) node.Node {
	return PreDec{
		"PreDec",
		variable,
	}
}

func (n PreDec) Name() string {
	return "PreDec"
}

func (n PreDec) Attributes() map[string]interface{} {
	return nil
}

func (n PreDec) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	v.LeaveNode(n)
}
