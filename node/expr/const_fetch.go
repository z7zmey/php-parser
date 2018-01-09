package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ConstFetch struct {
	Constant node.Node
}

func NewConstFetch(Constant node.Node) *ConstFetch {
	return &ConstFetch{
		Constant,
	}
}

func (n *ConstFetch) Attributes() map[string]interface{} {
	return nil
}

func (n *ConstFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Constant != nil {
		vv := v.GetChildrenVisitor("Constant")
		n.Constant.Walk(vv)
	}

	v.LeaveNode(n)
}
