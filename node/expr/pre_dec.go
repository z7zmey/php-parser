package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PreDec struct {
	Variable node.Node
}

func NewPreDec(Variable node.Node) *PreDec {
	return &PreDec{
		Variable,
	}
}

func (n *PreDec) Attributes() map[string]interface{} {
	return nil
}

func (n *PreDec) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
