package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ClassConstFetch struct {
	Class        node.Node
	ConstantName node.Node
}

func NewClassConstFetch(Class node.Node, ConstantName node.Node) *ClassConstFetch {
	return &ClassConstFetch{
		Class,
		ConstantName,
	}
}

func (n *ClassConstFetch) Attributes() map[string]interface{} {
	return nil
}

func (n *ClassConstFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ConstantName != nil {
		vv := v.GetChildrenVisitor("ConstantName")
		n.ConstantName.Walk(vv)
	}

	if n.Class != nil {
		vv := v.GetChildrenVisitor("Class")
		n.Class.Walk(vv)
	}

	v.LeaveNode(n)
}
