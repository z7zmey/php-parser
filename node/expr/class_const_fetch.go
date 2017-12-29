package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ClassConstFetch) Attributes() map[string]interface{} {
	return nil
}

func (n ClassConstFetch) Name() string {
	return "ClassConstFetch"
}

type ClassConstFetch struct {
	name         string
	class        node.Node
	constantName node.Node
}

func NewClassConstFetch(class node.Node, constantName node.Node) node.Node {
	return ClassConstFetch{
		"ClassConstFetch",
		class,
		constantName,
	}
}

func (n ClassConstFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.constantName != nil {
		vv := v.GetChildrenVisitor("constantName")
		n.constantName.Walk(vv)
	}

	if n.class != nil {
		vv := v.GetChildrenVisitor("class")
		n.class.Walk(vv)
	}

	v.LeaveNode(n)
}
