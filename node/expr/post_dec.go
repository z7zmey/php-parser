package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PostDec struct {
	Variable node.Node
}

func NewPostDec(Variable node.Node) *PostDec {
	return &PostDec{
		Variable,
	}
}

func (n *PostDec) Attributes() map[string]interface{} {
	return nil
}

func (n *PostDec) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
