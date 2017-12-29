package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ConstFetch struct {
	name     string
	constant node.Node
}

func NewConstFetch(constant node.Node) node.Node {
	return ConstFetch{
		"ConstFetch",
		constant,
	}
}

func (n ConstFetch) Name() string {
	return "ConstFetch"
}

func (n ConstFetch) Attributes() map[string]interface{} {
	return nil
}

func (n ConstFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.constant != nil {
		vv := v.GetChildrenVisitor("constant")
		n.constant.Walk(vv)
	}

	v.LeaveNode(n)
}
