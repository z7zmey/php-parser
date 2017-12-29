package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ConstList struct {
	name   string
	consts []node.Node
}

func NewConstList(consts []node.Node) node.Node {
	return ConstList{
		"ConstList",
		consts,
	}
}

func (n ConstList) Name() string {
	return "ConstList"
}

func (n ConstList) Attributes() map[string]interface{} {
	return nil
}

func (n ConstList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.consts != nil {
		vv := v.GetChildrenVisitor("consts")
		for _, nn := range n.consts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
