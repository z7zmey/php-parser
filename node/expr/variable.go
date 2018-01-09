package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Variable struct {
	VarName node.Node
}

func NewVariable(VarName node.Node) *Variable {
	return &Variable{
		VarName,
	}
}

func (n *Variable) Attributes() map[string]interface{} {
	return nil
}

func (n *Variable) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.VarName != nil {
		vv := v.GetChildrenVisitor("VarName")
		n.VarName.Walk(vv)
	}

	v.LeaveNode(n)
}
