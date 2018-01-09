package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ConstList struct {
	Consts []node.Node
}

func NewConstList(Consts []node.Node) *ConstList {
	return &ConstList{
		Consts,
	}
}

func (n *ConstList) Attributes() map[string]interface{} {
	return nil
}

func (n *ConstList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Consts != nil {
		vv := v.GetChildrenVisitor("Consts")
		for _, nn := range n.Consts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
