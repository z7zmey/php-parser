package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PreInc struct {
	Variable node.Node
}

func NewPreInc(Variable node.Node) *PreInc {
	return &PreInc{
		Variable,
	}
}

func (n *PreInc) Attributes() map[string]interface{} {
	return nil
}

func (n *PreInc) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
