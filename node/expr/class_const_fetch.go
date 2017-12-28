package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n ClassConstFetch) Name() string {
	return "ClassConstFetch"
}

type ClassConstFetch struct {
	name     string
	class    node.Node
	constant token.Token
}

// TODO: constant must be identifier
func NewClassConstFetch(class node.Node, constant token.Token) node.Node {
	return ClassConstFetch{
		"ClassConstFetch",
		class,
		constant,
	}
}

func (n ClassConstFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.Scalar("constant", n.constant.Value)

	if n.class != nil {
		vv := v.GetChildrenVisitor("class")
		n.class.Walk(vv)
	}

	v.LeaveNode(n)
}
