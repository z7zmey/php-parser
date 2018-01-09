package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PostInc struct {
	Variable node.Node
}

func NewPostInc(Variable node.Node) *PostInc {
	return &PostInc{
		Variable,
	}
}

func (n *PostInc) Attributes() map[string]interface{} {
	return nil
}

func (n *PostInc) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
